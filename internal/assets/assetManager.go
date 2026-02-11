package asset

import (
	"sync"
	"sync/atomic"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/google/uuid"
)



type AssetStore struct {
	asset ports.Asset
	referenceCount atomic.Int32
	assetMutex sync.Mutex
}

type AssetManager struct {
	storeList types.TypeSyncMap[uuid.UUID,*AssetStore]
	factoryList types.TypeSyncMap[uuid.UUID,func() ports.Asset] 
}


var manager *AssetManager
var managerOnce sync.Once

func GetAssetManager() *AssetManager{
	managerOnce.Do(func() {
		manager = &AssetManager{
			storeList: types.TypeSyncMap[uuid.UUID,*AssetStore] {},
			factoryList: types.TypeSyncMap[uuid.UUID,func() ports.Asset] {},
		}
		manager.Init()
	})
	return manager
}

func newAssetStore(asset ports.Asset) *AssetStore {
	store := &AssetStore{
		asset: asset,
		referenceCount: atomic.Int32{},
	}
	store.referenceCount.Store(0)
	return store
}


func (instance *AssetManager) get(uuid uuid.UUID) (*AssetStore,error) {
 	store, exist := instance.storeList.Load(uuid)
	if !exist {
		asset := instance.createAsset(uuid)
		if asset == nil {
			return nil , packagederror.NewError(packagederror.UnknownAsset,"FailToLoadAsset")
		}
		newStore := newAssetStore(asset)
		store , _ = instance.storeList.LoadOrStore(uuid,newStore)
	}
	count := store.referenceCount.Add(1)
	
	if count == 1 {
		store.assetMutex.Lock()
		if !store.asset.IsLoaded() {
			if err := store.asset.Load() ; err != nil {
				store.assetMutex.Unlock()
				return nil, err
			}
		}
		store.assetMutex.Unlock()
	}
	return store , nil
}

func (instance *AssetManager) createAsset(uuid uuid.UUID) ports.Asset {
	create , ok := instance.factoryList.Load(uuid)
	if !ok {
		return nil
	}
	return create()
}

func (instance *AssetManager) register(uuid uuid.UUID,createFunc func() ports.Asset) error {
	if _ , ok := instance.factoryList.Load(uuid) ; ok {
		return packagederror.NewError(packagederror.FailRegisterDuplicate,"Already registered")
	}
	instance.factoryList.Store(uuid,createFunc)
	return nil
}

func (instance *AssetManager) Release(uuid uuid.UUID) {
	store , ok := instance.storeList.Load(uuid)
	if !ok {
		return
	}
	count := store.referenceCount.Add(-1)
	if count == 0 {
		store.assetMutex.Lock()
		if store.referenceCount.Load() == 0 {
			store.asset.UnLoad()	
			instance.storeList.Delete(uuid)
		}
		store.assetMutex.Unlock()
	}
}

func (instance *AssetManager) ShaderAsset(uuid uuid.UUID) (*types.Reference[types.Program], error) {
	store ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	shader ,ok := store.asset.(*ShaderAsset)
	if !ok {
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return types.NewReference(shader.Get(),func(){ instance.Release(uuid) }) , nil 
}

func (instance *AssetManager) TextureAsset(uuid uuid.UUID) (*types.Reference[types.Texture], error) {
	store ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	texture ,ok := store.asset.(*TextureAsset)
	if !ok {
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return types.NewReference(texture.Get(),func(){ instance.Release(uuid) }) , nil 
}
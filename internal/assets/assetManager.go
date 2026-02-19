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
	IsLoad atomic.Bool
	assetMutex sync.Mutex
}

type AssetManager struct {
	storeList map[uuid.UUID]*AssetStore //types.TypeSyncMap[uuid.UUID,*AssetStore]
	ListMutex sync.RWMutex
	factoryList types.TypeSyncMap[uuid.UUID,func() ports.Asset] 
}


var manager *AssetManager
var managerOnce sync.Once

func GetAssetManager() *AssetManager{
	managerOnce.Do(func() {
		manager = &AssetManager{
			storeList: make(map[uuid.UUID]*AssetStore), //types.TypeSyncMap[uuid.UUID,*AssetStore] {},
			factoryList: types.TypeSyncMap[uuid.UUID,func() ports.Asset] {},
		}
	})
	return manager
}

func newAssetStore(asset ports.Asset) *AssetStore {
	store := &AssetStore{
		asset: asset,
		referenceCount: atomic.Int32{},
		IsLoad: atomic.Bool{},
	}
	store.IsLoad.Store(false)
	store.referenceCount.Store(0)
	return store
}


func (instance *AssetManager) get(uuid uuid.UUID) (*AssetStore,error) {
 	instance.ListMutex.Lock()
	store , exist := instance.storeList[uuid]
	if !exist {
		asset := instance.createAsset(uuid)
		if asset == nil {
			instance.ListMutex.Unlock()
			return nil , packagederror.NewError(packagederror.UnknownAsset,"FailToLoadAsset")
		}
		newStore := newAssetStore(asset)
		instance.storeList[uuid] = newStore
		store = newStore
	}
	store.referenceCount.Add(1)

	if store.IsLoad.Load() {
		instance.ListMutex.Unlock()
		return store , nil
	}

	instance.ListMutex.Unlock()
	return instance.storeLoad(store,uuid)
}

func (instance *AssetManager) storeLoad(store *AssetStore,uuid uuid.UUID) (*AssetStore,error) {
	store.assetMutex.Lock()
	defer store.assetMutex.Unlock()

	if store.asset.IsLoaded() {
		return store , nil
	}

	err := store.asset.Load() 

	if err == nil {
		store.IsLoad.Store(true)
		return store , nil
	}

	instance.ListMutex.Lock()
	if count := store.referenceCount.Add(-1) ; count == 0 {
		delete(instance.storeList,uuid)
	}
	instance.ListMutex.Unlock()
	return  nil , err
} 

func (instance *AssetManager) createAsset(uuid uuid.UUID) ports.Asset {
	create , ok := instance.factoryList.Load(uuid)
	if !ok {
		return nil
	}
	return create()
}

func (instance *AssetManager) Register(uuid uuid.UUID,createFunc func() ports.Asset) error {
	if _ , ok := instance.factoryList.Load(uuid) ; ok {
		return packagederror.NewError(packagederror.FailRegisterDuplicate,"Already registered")
	}
	instance.factoryList.Store(uuid,createFunc)
	return nil
}

func (instance *AssetManager) Release(uuid uuid.UUID) {
	instance.ListMutex.Lock()
	store , ok := instance.storeList[uuid]
	if !ok {
		instance.ListMutex.Unlock()
		return
	}
	store.referenceCount.Add(-1)
	store.assetMutex.Lock()

	if store.referenceCount.Load() != 0 {
		instance.ListMutex.Unlock()
		store.assetMutex.Unlock()
		return
	}

	delete(instance.storeList,uuid)
	instance.ListMutex.Unlock()

	store.asset.UnLoad()
	store.assetMutex.Unlock()
}

func (instance *AssetManager) ShaderAsset(uuid uuid.UUID) (*types.Reference[types.Program], error) {
	store ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	shader ,ok := store.asset.(*ShaderAsset)
	if !ok {
		instance.Release(uuid)
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
		instance.Release(uuid)
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return types.NewReference(texture.Get(),func(){ instance.Release(uuid) }) , nil 
}

func (instance *AssetManager) MeshAsset2D(uuid uuid.UUID) (*types.Reference[types.Mesh], error) {
	store ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	mesh ,ok := store.asset.(*MeshAsset2D)
	if !ok {
		instance.Release(uuid)
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return types.NewReference(mesh.Get(),func(){ instance.Release(uuid) }) , nil 
}

func (instance *AssetManager) Material(uuid uuid.UUID) (*types.Handle[types.Material] , error) {
	store , err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	material , ok := store.asset.(*MaterialAsset)
	if !ok {
		instance.Release(uuid)
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return types.NewHandle(material.Get(),func() { instance.Release(uuid) },material.GetRenderer()) , nil
}
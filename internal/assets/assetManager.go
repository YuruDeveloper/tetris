package asset

import (
	"sync"

	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/google/uuid"
)


var manager *AssetManager
var managerOnce sync.Once

func GetAssetManager() *AssetManager{
	managerOnce.Do(func() {
		manager = &AssetManager{
			assetList: make(map[uuid.UUID]ports.Asset),
			referenceCount: make(map[uuid.UUID]int),
			factoryList: make(map[uuid.UUID]func() ports.Asset),
		}
		manager.Init()
	})
	return manager
}

type AssetManager struct {
	assetList map[uuid.UUID]ports.Asset
	referenceCount map[uuid.UUID]int
	factoryList map[uuid.UUID]func() ports.Asset
	mutex sync.RWMutex
}

func (instance *AssetManager) get(uuid uuid.UUID) (ports.Asset,error) {
	instance.mutex.Lock()
	defer instance.mutex.Unlock()
 	asset, exist := instance.assetList[uuid]
	if !exist {
		asset = instance.createAsset(uuid)
		if asset == nil {
			return nil , packagederror.NewError(packagederror.UnknownAsset,"FailToLoadAsset")
		}
		instance.assetList[uuid] = asset
		instance.referenceCount[uuid] = 0
	}
	if asset.IsLoaded() {
		instance.referenceCount[uuid] += 1
		return asset , nil
	}
	if err := asset.Load() ; err != nil {
		return nil , err
	}
	instance.referenceCount[uuid] += 1
	return asset , nil
}

func (instance *AssetManager) createAsset(uuid uuid.UUID) ports.Asset {
	create , ok := instance.factoryList[uuid]
	if !ok {
		return nil
	}
	return create()
}

func (instance *AssetManager) register(uuid uuid.UUID,createFunc func() ports.Asset) error {
	instance.mutex.Lock()
	defer instance.mutex.Unlock()
	if _ , ok := instance.factoryList[uuid] ; ok {
		return packagederror.NewError(packagederror.FailRegisterDuplicate,"Alreay regitryed factory")
	}
	instance.factoryList[uuid] = createFunc
	return nil
}

func (instance *AssetManager) Release(uuid uuid.UUID) {
	instance.mutex.Lock()
	defer instance.mutex.Unlock()
	_ , ok := instance.referenceCount[uuid]
	if !ok {
		return
	}
	instance.referenceCount[uuid] -= 1
	if instance.referenceCount[uuid] == 0 {
		instance.assetList[uuid].UnLoad()
		delete(instance.assetList,uuid)
		delete(instance.referenceCount,uuid)
	}
}

func (instance *AssetManager) ShaderAsset(uuid uuid.UUID) (*ShaderAsset, error) {
	asset ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	shader ,ok := asset.(*ShaderAsset)
	if !ok {
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return shader , nil
}

func (instance *AssetManager) TextureAsset(uuid uuid.UUID) (*TextureAsset, error) {
	asset ,err := instance.get(uuid)
	if err != nil {
		return nil , err
	}
	texture ,ok := asset.(*TextureAsset)
	if !ok {
		return nil , packagederror.NewError(packagederror.FailAssetTypeConvert,"Fail To Convert Asset Type")
	}
	return texture , nil
}
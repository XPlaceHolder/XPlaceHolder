package commands

import (
	"fmt"
	"os"

	builtinmanifests "github.com/kun-lun/artifacts/pkg/apis/builtinmanifests"
	"github.com/kun-lun/common/fileio"
	"github.com/kun-lun/common/flags"
	"github.com/kun-lun/common/helpers"
	"github.com/kun-lun/common/storage"
	digester "github.com/kun-lun/digester/pkg/apis"
)

type Digest struct {
	stateStore   storage.Store
	envIDManager helpers.EnvIDManager

	fs fileio.Fs
}

type DiegestConfig struct {
	Name string
}

func NewDigest(
	stateStore storage.Store,
	envIDManager helpers.EnvIDManager,
	fs fileio.Fs,
) Digest {
	return Digest{
		stateStore:   stateStore,
		envIDManager: envIDManager,
		fs:           fs,
	}
}

func (p Digest) CheckFastFails(args []string, state storage.State) error {
	config, err := p.ParseArgs(args, state)
	if err != nil {
		return err
	}
	if state.EnvID != "" && config.Name != "" && config.Name != state.EnvID {
		return fmt.Errorf("The env name cannot be changed for an existing environment. Current name is %s", state.EnvID)
	}
	return nil
}

func (p Digest) ParseArgs(args []string, state storage.State) (DiegestConfig, error) {
	var (
		config DiegestConfig
	)

	digestFlags := flags.New("digest")
	digestFlags.String(&config.Name, "name", os.Getenv("KL_ENV_NAME"))

	err := digestFlags.Parse(args)
	if err != nil {
		return DiegestConfig{}, err
	}
	return config, nil
}

func (p Digest) Execute(args []string, state storage.State) error {
	config, err := p.ParseArgs(args, state)
	if err != nil {
		return err
	}
	_, err = p.initialize(config, state)

	// choose one manifest from the artifacts galary.

	return err
}

func (p Digest) initialize(config DiegestConfig, state storage.State) (storage.State, error) {
	var err error
	state, err = p.envIDManager.Sync(state, config.Name)
	if err != nil {
		return storage.State{}, fmt.Errorf("Env id manager sync: %s", err)
	}

	err = p.stateStore.Set(state)
	if err != nil {
		return storage.State{}, fmt.Errorf("Save state: %s", err)
	}

	questionaireFilePath, err := p.stateStore.GetQuestionaireFilePath()
	if err := digester.Run(questionaireFilePath); err != nil {
		return storage.State{}, fmt.Errorf("Call digester: %s", err)
	}

	err = p.pickUpManifest()

	return state, err
}

// TODO(zhongyi) This is a stub, should pick the manifest up based on the q&a file.
func (p Digest) pickUpManifest() error {
	var (
		isSmallPHP bool
	)
	isSmallPHP = true
	artifactFilePath, err := p.stateStore.GetMainArtifactFilePath()
	if err != nil {
		return err
	}
	if isSmallPHP {
		content, err := builtinmanifests.FSByte(false, "/manifests/small_php.yml")
		if err != nil {
			return err
		}
		err = p.fs.WriteFile(artifactFilePath, content, 0644)
		return err
	}
	return fmt.Errorf("we only support the small php")
}

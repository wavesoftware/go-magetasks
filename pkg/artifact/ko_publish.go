package artifact

import (
	"fmt"

	"github.com/google/ko/pkg/build"
	"github.com/google/ko/pkg/commands"
	"github.com/google/ko/pkg/commands/options"
	"github.com/wavesoftware/go-magetasks/config"
)

const koPublishResult = "ko.publish.result"

// KoPublisher publishes images with Google's KO.
type KoPublisher struct{}

func (kp KoPublisher) Accepts(artifact config.Artifact) bool {
	_, ok := artifact.(Image)
	return ok
}

func (kp KoPublisher) Publish(artifact config.Artifact, notifier config.Notifier) config.Result {
	image, ok := artifact.(Image)
	if !ok {
		return config.Result{Error: ErrInvalidArtifact}
	}
	buildResult, ok := config.Actual().Context.Value(BuildKey(image)).(config.Result)
	if !ok || buildResult.Failed() {
		return config.Result{Error: fmt.Errorf(
			"%w: can't find successful KO build result", ErrInvalidArtifact)}
	}
	result, ok := buildResult.Info[koBuildResult].(build.Result)
	if !ok {
		return config.Result{Error: fmt.Errorf(
			"%w: can't find successful KO build result", ErrInvalidArtifact)}
	}
	po := &options.PublishOptions{}
	ctx := config.Actual().Context
	publisher, err := commands.NewPublisher(po)
	if err != nil {
		return resultErrKoFailed(err)
	}
	ref, err := publisher.Publish(ctx, result, image.GetName())
	if err != nil {
		return resultErrKoFailed(err)
	}
	return config.Result{Info: map[string]interface{}{
		koPublishResult: ref,
	}}
}

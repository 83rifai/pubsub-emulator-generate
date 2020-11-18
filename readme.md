## run pubsub emulator

export PUBSUB_PROJECT_ID=ctcd-omnichannel-dev && export PUBSUB_EMULATOR_HOST=localhost:8085
gcloud beta emulators pubsub start --project=ctcd-omnichannel-dev

export PUBSUB_PROJECT_ID=ctcd-omnichannel-stg && export PUBSUB_EMULATOR_HOST=localhost:8085
gcloud beta emulators pubsub start --project=ctcd-omnichannel-stg

export PUBSUB_PROJECT_ID=tri-omnichannel-prod && export PUBSUB_EMULATOR_HOST=localhost:8085
gcloud beta emulators pubsub start --project=tri-omnichannel-prod

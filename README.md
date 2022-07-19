# SAM AWS Go Start + Task File

## Prerequisits

- Install task from https://task.dev
- Install Docker
- Install AWS CLI
- Install SAM CLI

## Description

`task add -- <NAME>` creates a new dir in the lambda folder with a main.go, go.mod, go.sum which contains the necessary lambda packages.

`task start` runs sam build + sam local start-api with the warm containers flag set to eager, otherwise api calls take 2s+

currently `task add` does not update the template.yaml file to include the new endpoints, that has to be done manually.

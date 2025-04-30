#!/bin/bash

# Map of keywords to actual commands
declare -A COMMAND_MAP

# Define your mappings here
COMMAND_MAP=(
  ["air"]="go run github.com/air-verse/air@latest"
)

# Ensure a command key is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <command_key> [flags...]"
  echo "Available commands:"
  for key in "${!COMMAND_MAP[@]}"; do
    echo "  $key -> ${COMMAND_MAP[$key]}"
  done
  exit 1
fi

COMMAND_KEY="$1"
shift # Remove the first argument so that $@ now contains only flags

COMMAND="${COMMAND_MAP[$COMMAND_KEY]}"

if [ -z "$COMMAND" ]; then
  echo "Error: Unknown command key '$COMMAND_KEY'"
  echo "Available commands:"
  for key in "${!COMMAND_MAP[@]}"; do
    echo "  $key -> ${COMMAND_MAP[$key]}"
  done
  exit 1
fi

# Construct full command with flags
FULL_COMMAND="$COMMAND $*"

# Print and run the command
echo "Running: $FULL_COMMAND"
eval $FULL_COMMAND

#!/usr/bin/env bash
set -e

oapi-codegen --config api/types.cfg.yml api/main.all.yml

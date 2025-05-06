#!/bin/bash

# Delve (dlv) をインストール
echo "Installing Delve debugger..."
go install github.com/go-delve/delve/cmd/dlv@latest

echo "Delve installation completed."

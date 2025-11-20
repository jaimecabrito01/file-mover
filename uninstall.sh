#!/bin/bash

set -e

SERVICE_NAME="filemover.service"
SERVICE_PATH="/etc/systemd/system/$SERVICE_NAME"
BIN_PATH="/usr/local/bin/filemover"
INSTALL_DIR="/opt/filemover"

echo "=== Desinstalando FileMover ==="

# Parar o serviço
if systemctl list-units --full -all | grep -q "$SERVICE_NAME"; then
  echo "Parando serviço..."
  sudo systemctl stop "$SERVICE_NAME" || true
  echo "Desabilitando serviço..."
  sudo systemctl disable "$SERVICE_NAME" || true
fi

# Remover arquivo do serviço
if [ -f "$SERVICE_PATH" ]; then
  echo "Removendo arquivo de serviço systemd: $SERVICE_PATH"
  sudo rm -f "$SERVICE_PATH"
  echo "Recarregando systemd..."
  sudo systemctl daemon-reload
fi

# Remover binário
if [ -f "$BIN_PATH" ]; then
  echo "Removendo binário: $BIN_PATH"
  sudo rm -f "$BIN_PATH"
else
  echo "Binário não encontrado em $BIN_PATH"
fi

# Remover diretório de instalação
if [ -d "$INSTALL_DIR" ]; then
  echo "Removendo diretório da aplicação: $INSTALL_DIR"
  sudo rm -rf "$INSTALL_DIR"
else
  echo "Diretório $INSTALL_DIR não encontrado."
fi

echo "FileMover desinstalado com sucesso."

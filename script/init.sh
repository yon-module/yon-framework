#!/bin/bash

# **1. Menampilkan versi dan informasi**
SCRIPT_VERSION="1.0.0"
DEFAULT_PACKAGE_FRAMEWORK=github.com/yon-module/yon-framework
echo "======================================="
echo " Golang Project Setup Script v$SCRIPT_VERSION"
echo "======================================="

# **2. Input Project Name**
read -p "Enter Project Name: " PROJECT_NAME

# Ubah ke kebab-case (lowercase dan ganti spasi dengan '-')
FOLDER_NAME=$(echo "$PROJECT_NAME" | tr '[:upper:]' '[:lower:]' | tr ' ' '-')
echo "Project folder name: $FOLDER_NAME"

# **3. Input Package Name (default = folder name)**
read -p "Enter Package Name (default: $FOLDER_NAME): " PACKAGE_NAME
PACKAGE_NAME=${PACKAGE_NAME:-$FOLDER_NAME}

# **4. Generate Folder Structure**
echo "Creating project structure..."
mkdir -p $FOLDER_NAME/src/{main/{controllers,model/{dto/{request,response},entity},repository,service,util,helper},tests}
touch $FOLDER_NAME/.env.sample
touch $FOLDER_NAME/main.go

# **5. Create .env.sample**
echo "Setup env sample..."
cat <<EOL > $FOLDER_NAME/.env.sample
yon.server.appName=$PACKAGE_NAME
yon.server.port=8080

yon.database.type=
yon.database.host=
yon.database.user=
yon.database.password=
yon.database.port=
yon.database.db=
EOL

# **6. Run install Go packages**
echo "Installing required Go packages..."
cd $FOLDER_NAME
go mod init $PACKAGE_NAME
go get -u $DEFAULT_PACKAGE_FRAMEWORK

# Selesai
echo "Project setup complete! 🎉"
echo "Navigate to your project folder: cd $FOLDER_NAME"

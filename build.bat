@echo off
echo "start build"
fyne package -icon ./res-static/icon/icon.png

echo "copy files to build-target"
MD build-target\
MD build-target\res-static
MD build-target\res-static\fonts
MD build-target\res-static\icon

echo "move build target..."
MOVE ScrcpyLauncher.exe build-target\
echo ""

echo "copy resource"
COPY res-static\fonts\*.* build-target\res-static\fonts\
COPY res-static\icon\*.* build-target\res-static\icon\
echo ""


echo "build success, you can open binary file in build-target"
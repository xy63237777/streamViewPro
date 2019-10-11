#! /bin/bash
# Build Web UI

cd ~/go/src/streamViewPro/web
go install
cp ~/go/bin/web ~/go/bin/video_server_web_ui/web
cp -R ~/go/src/streamViewPro/template/ ~/go/bin/video_server_web_ui/

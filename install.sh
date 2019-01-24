#!/usr/bin/env bash
FILE="/lib/systemd/system/geoip-api.service"

geoipFile="geoip-api-linux-amd64"
curl -s https://api.github.com/repos/danielcherubini/geoip-api/releases | grep browser_download_url | grep ${geoipFile} | head -n 1 | cut -d '"' -f 4 | wget -i -


chmod +x ${geoipFile}
mv ${geoipFile} /usr/local/bin/geoip-api

cat > $FILE <<- EOM
[Unit]
Description=GEO IP to Locale

[Service]
Type=simple
PIDFile=/tmp/geoip-api.pid-4040
Restart=always
ExecStart=/usr/local/bin/geoip-api -lang=/etc/geoip-api/languages.json

[Install]
WantedBy=multiuser.target
EOM


file="/etc/geoip-api/languages.json"
if [ -f "$file" ]
then
	echo "$file found, skipping."
else
	# echo "$file not found, adding."
    # sudo mkdir /etc/geoip-api
    # sudo touch /etc/geoip-api/languages.json
    # sudo chmod 777 /etc/geoip-api/languages.json
	exit 1
fi

sudo systemctl daemon-reload
sudo systemctl start geoip-api.service

curl -O .zip
unzip pulse.zip -d /usr/local/
grep -q "PULSE_HOME" ~/.bashrc || append_env

function append_env() {
  cat <<EOF >>~/.bashrc
export PULSE_HOME=/usr/local/pulse
export PATH=$PATH:$PULSE_HOME/bin
export PULSE_CONFIG_FILE=$PULSE_HOME/config.yaml
export PULSE_REPORT_FILE=$PULSE_HOME/report.tpl
EOF
}

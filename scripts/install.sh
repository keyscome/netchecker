curl -O .zip
unzip netchecker.zip -d /usr/local/
grep -q "NETCHECKER_HOME" ~/.bashrc || append_env

function append_env()
{
  cat <<EOF >> ~/.bashrc
export NETCHECKER_HOME=/usr/local/netchecker
export PATH=$PATH:$NETCHECKER_HOME/bin
export NETCHECKER_CONFIG_FILE=$NETCHECKER_HOME/config.yaml
export NETCHECKER_REPORT_FILE=$NETCHECKER_HOME/report.tpl
EOF
}

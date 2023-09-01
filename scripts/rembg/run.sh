WORK_DIR=$(pwd)
VENV_DIR=$WORK_DIR/scripts/rembg/venv

source $VENV_DIR/bin/activate

$VENV_DIR/bin/python $WORK_DIR/scripts/rembg/main.py -i $1 -o $2

if [ $? -eq 0 ]; then
	echo "rembg script succeeded"
	exit 0
else
	echo "rembg script failed" >&2
	exit 1
fi

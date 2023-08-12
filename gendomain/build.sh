# !/bin/sh

# ROOTPATH=`pwd`
ROOTPATH="."

# Module Name
MODULE_NAME=`head -n 1 ../go.mod`
MODULE_NAME=${MODULE_NAME/module /}

# Build Information
BUILD_VERSION=`cat version.txt`
BUILD_TIME=`date "+%F %T"`
BUILD_NAME_SUFFIX=`date "+%Y%m%d%H"`
COMMIT_SHA1=`git rev-parse HEAD`

echo BUILD_NAME_SUFFIX=${BUILD_NAME_SUFFIX}
echo BUILD_TIME=${BUILD_TIME}
echo BUILD_VERSION=${BUILD_VERSION}
echo COMMIT_SHA1=${COMMIT_SHA1}

tag=`date +'%Y-%m-%d-%H%M%S'`

# Target Folder
mkdir -p ${ROOTPATH}/bin
rm -rf ${ROOTPATH}/bin/$1*

# System Information
go env | egrep "GOARCH|GOOS|GOHOSTARCH|GOHOSTOS"

# Project List
projects=("gendomain")

function build()
{
  PROJ=$1
  NAME=$2
  REPLACE_CMD=$3

  # cd ${ROOTPATH}/cmd/${PROJ} && \
  go build \
    -ldflags " \
    -X '${MODULE_NAME}/src/config.BuildVersion=${BUILD_VERSION}' \
    -X '${MODULE_NAME}/src/config.BuildName=${PROJ}' \
    -X '${MODULE_NAME}/src/config.BuildTime=${BUILD_TIME}' \
    -X '${MODULE_NAME}/src/config.BuildCommitID=${COMMIT_SHA1}' \
    -X '${MODULE_NAME}/src/syscmd.REPLACE_CMD=${REPLACE_CMD}' \
    " \
    -o "${ROOTPATH}/bin/${NAME}" .

  chmod +x ${ROOTPATH}/bin/${NAME}
}

for proj in "${projects[@]}";
do
  if [ "$1" == "" ]; then
    # build ${proj} ${proj//find/replace};
    echo "building ${proj}"
    # build ${proj} ${proj}-${tag}
    build ${proj} ${proj} ${proj}
  elif [ "$1" == "${proj}" ]; then
    # build ${proj} ${proj//find/replace};
    echo "building ${proj}"
    # build ${proj} ${proj}-${tag}
    build ${proj} ${proj} ${proj}
  fi
done

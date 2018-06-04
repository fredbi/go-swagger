#! /bin/bash 
if [[ ${1} == "--clean" ]] ; then
    clean=1
fi
# Acknowledgements
#testcases="${testcases} ../../codegen/todolist.models.yml"   
#testcases="${testcases} ../../codegen/todolist.discriminators.yml"   
#testcases="${testcases} fixture-sanity.yaml"   
testcases="${testcases} ../1409/fixture-1409.yaml"   
#testcases="${testcases} ../1341/swagger.yaml"   
testcases="${testcases} fixture-1232.yaml"   
for testcase in ${testcases} ; do
    spec=${testcase}
    testcase=`basename ${testcase}`
    target=./gen-${testcase%.*}
    serverName="codegensrv"
    rm -rf ${target}
    mkdir ${target}
    echo "Server generation for ${spec}"
    swagger generate server --skip-validation --spec=${spec} --target=${target} --name=${serverName} --output=${testcase%.*}.log
    if [[ $? != 0 ]] ; then
        echo "Generation failed for ${spec}"
        exit 1
    fi
    echo "${spec}: Generation OK"
    if [[ ! -d ${target}/models ]] ; then
        echo "No model in this spec!"
    else
        (cd ${target}/cmd/${serverName}-server; go build)
        if [[ $? != 0 ]] ; then
            echo "Build failed for ${spec}"
            exit 1
        fi
        echo "${spec}: Build OK"
        if [[ -n ${clean} ]] ; then 
             rm -rf ${target}
        fi
    fi
done

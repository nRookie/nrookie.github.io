

https://docs.hpc.sjtu.edu.cn/app/ai/pytorch.html



``` shell
conda create -n pytorch-env -y
conda activate pytorch-env
conda install -c conda-forge matplotlib  pandas scipy -y
```





``` shell
cd ../epc-task-au0tqabx
echo 'create pytorch env'
conda create -n pytorch-env -y
echo 'activate pytorch-env'
conda activate pytorch-env
echo 'init bash env'
conda init bash
echo 'start a new bash env'
exec bash
echo $SHELL
bash
echo 'source bash'
source ~/.bashrc 
echo 'activate pytorch-env'
conda activate pytorch-env
conda install -c conda-forge matplotlib  pandas scipy -y
python3 main.py
exit 1
```


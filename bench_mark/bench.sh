#!/bin/bash

time python ./pazdora-cal_original.py > banch_result.txt
time python ./pazdora-cal_チェックロジックを修正した_1.py > banch_result.txt
time python ./pazdora-cal_盤面生成を最初だけにした_2.py > banch_result.txt
time python ./pazdora-cal_wotsushiによる改善_3.py > banch_result.txt
time python ./pazdora-cal_ドロップの数のみを格納したやつ_4.py > banch_result.txt
time python ./pazdora-cal_parallel_5.py > banch_result.txt

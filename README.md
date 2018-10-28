# pazdora-cal
パズ◯ラの盤面の欠損率を計算するスクリプトです。

goバージョンが一番早いです

## goバージョンの使い方

```
$ cd go
$ ./pazdora-cal [option]
```

## option 


* -x [int] : 盤面の縦方向の大きさを指定
* -y [int] : 盤面の縦方向の大きさを指定
* -loopCnt [int] : ループ回数を指定


例

```
$ ./pazdora-cal -x 7 -y 6 -loopCnt 1000000
```

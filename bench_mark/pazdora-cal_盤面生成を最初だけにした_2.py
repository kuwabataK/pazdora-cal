from numpy.random import *
from itertools import chain

def check_normal_drops (drops):
    ## ドロップが3つ以上つながっているかどうかをチェックする
    for idx_i, val_i in enumerate(drops):
        for idx_j, val_j in enumerate(val_i):

            if idx_j + 2 < len(val_i):
                if val_j == val_i[idx_j + 1] and val_i[idx_j + 1] == val_i[idx_j + 2]:
                    return False
            
            if idx_i + 2 < len(drops):
                if val_j == drops[idx_i + 1][idx_j] and val_j == drops[idx_i + 2][idx_j]:
                    return False

    return True

def generate_drops(height, width):
    return randint(0,6,(height,width))

##============= カスタムフィールド==============##

## 試行回数を指定
loop_cnt = 1000000

## 盤面を指定
height = 5
width = 6

##=============================================##

## 指定した欠損条件に対して、モンテカルロ法により、欠損しなかった回数と欠損した回数の組を返す
def monte_carlo_freq(lack_cond,fields):
    num_ok = 0
    num_ng = 0
    for drops in fields:

        fire_drops_num = len([i for i in drops if i == 0])
        blue_drops_num = len([i for i in drops if i == 1])
        green_drops_num = len([i for i in drops if i == 2])
        light_drops_num = len([i for i in drops if i == 3])
        black_drops_num = len([i for i in drops if i == 4])
        recovery_drops_num = len([i for i in drops if i == 5])

        if not lack_cond(
                fire_drops_num,
                blue_drops_num,
                green_drops_num,
                light_drops_num,
                black_drops_num,
                recovery_drops_num):
            num_ok = num_ok +1
        else:
            num_ng = num_ng +1
    return num_ok, num_ng

def generate_fields():

    fields = []

    for x in range(loop_cnt):
        drops = generate_drops(height,width)
        if check_normal_drops(drops) == False:
            continue

        drops = list(chain.from_iterable(drops))
        fields.append(drops)
    return fields

## 試行回数、OKの回数、NGの回数、確率を出力する
def print_prob(num_ok, num_ng):
    print("試行回数は" + str(num_ok + num_ng))
    print("OKの回数は" + str(num_ok))
    print("NGの回数は" + str(num_ng))
    print("確率は" + str(num_ok/(num_ok + num_ng) * 100))


##=============================================##

drops = generate_drops(height,width)
while check_normal_drops(drops) == False:
    drops = generate_drops(height,width)
print(drops)

print("総試行回数は" + str(loop_cnt) + "回")

fields = generate_fields()

## 2色いずれか5個以上ある確率
print("=========================")
print("2色いずれか5個以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 4 and d <= 4, fields)
print_prob(num_ok, num_ng)

## 指定2色がある確率
print("=========================")
print("指定2色がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 2 or d <= 2, fields)
print_prob(num_ok, num_ng)

## 6色のうち5色がある確率
print("=========================")
print("6色のうち5色がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: sorted((f, b, g, l, d, r))[1] <= 2, fields)
print_prob(num_ok, num_ng)

## 5色のうち4色がある確率
print("=========================")
print("5色のうち4色がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: sorted((f, b, g, l, d))[1] <= 2, fields)
print_prob(num_ok, num_ng)

## 指定同色2コンボがある確率
print("=========================")
print("指定同色2コンボ(6個)がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 5, fields)
print_prob(num_ok, num_ng)

## 指定同色1コンボがある確率
print("=========================")
print("指定同色1コンボ(3個)がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 2, fields)
print_prob(num_ok, num_ng)

## 指定同色が5個以上がある確率

print("=========================")
print("指定同色が5個以上がある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 4, fields)
print_prob(num_ok, num_ng)

## 全色ある確率
print("=========================")
print("全色ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: sorted((f, b, g, l, d, r))[0] <= 2, fields)
print_prob(num_ok, num_ng)

## ガードブレイクが発動できる確率
print("=========================")
print("ガードブレイクが発動できる確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: sorted((f, b, g, l, d))[0] <= 2, fields)
print_prob(num_ok, num_ng)

## 指定一色が8個以上ある確率(ねね/覚醒カストル等)
print("=========================")
print("指定一色が8個以上ある確率(ねね/覚醒カストル等)")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 7, fields)
print_prob(num_ok, num_ng)

## 指定一色が4個以上ある確率(進化前リーチェ)
print("=========================")
print("指定一色が4個以上ある確率(進化前リーチェ)")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: f <= 3, fields)
print_prob(num_ok, num_ng)

## 7コンボ以上ある確率
print("=========================")
print("7コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 6, fields)
print_prob(num_ok, num_ng)

## 8コンボ以上ある確率
print("=========================")
print("8コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 7, fields)
print_prob(num_ok, num_ng)

## 9コンボ以上ある確率
print("=========================")
print("9コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 8, fields)
print_prob(num_ok, num_ng)

## 10コンボ以上ある確率
print("=========================")
print("10コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 9, fields)
print_prob(num_ok, num_ng)

## 11コンボ以上ある確率
print("=========================")
print("11コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 10, fields)
print_prob(num_ok, num_ng)

## 12コンボ以上ある確率
print("=========================")
print("12コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 11, fields)
print_prob(num_ok, num_ng)

## 13コンボ以上ある確率
print("=========================")
print("13コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 12, fields)
print_prob(num_ok, num_ng)

## 14コンボ以上ある確率
print("=========================")
print("14コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 3) + (b // 3) + (g // 3) + (l // 3) + (d // 3) + (r // 3) <= 13, fields)
print_prob(num_ok, num_ng)

## (4個消しで)5コンボ以上ある確率
print("=========================")
print("(4個消しで)5コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 4, fields)
print_prob(num_ok, num_ng)

## (4個消しで)6コンボ以上ある確率
print("=========================")
print("(4個消しで)6コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 5, fields)
print_prob(num_ok, num_ng)

## (4個消しで)7コンボ以上ある確率
print("=========================")
print("(4個消しで)7コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 6, fields)
print_prob(num_ok, num_ng)

## (4個消しで)8コンボ以上ある確率
print("=========================")
print("(4個消しで)8コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 7, fields)
print_prob(num_ok, num_ng)

## (4個消しで)9コンボ以上ある確率
print("=========================")
print("(4個消しで)9コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 8, fields)
print_prob(num_ok, num_ng)

## (4個消しで)10コンボ以上ある確率
print("=========================")
print("(4個消しで)10コンボ以上ある確率")
num_ok, num_ng = monte_carlo_freq(lambda f, b, g, l, d, r: (f // 4) + (b // 4) + (g // 4) + (l // 4) + (d // 4) + (r // 4) <= 9, fields)
print_prob(num_ok, num_ng)
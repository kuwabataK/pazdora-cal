##============= カスタムフィールド==============##

## 試行回数を指定
const loop_cnt = 100000
## 盤面を指定
const height = 7
const width = 6

##=============================================##

function main()
    function check_normal_drops(drops)
        ## ドロップが3つ以上つながっているかどうかをチェックする
        for (idx_i, val_i) in enumerate(drops)
            for (idx_j, val_j) in enumerate(val_i)
                if val_j == val_i[idx_j + 1] == val_i[idx_j + 2]
                    return false
                end
                if idx_j + 3 >= length(val_i)
                    break
                end
            end
        end
        ## 配列を逆転して再度計算
        drops = [[drops[j][i] for j in 1:length(drops)] for i in 1:length(drops[1])]
        for (idx_i, val_i) in enumerate(drops)
            for (idx_j, val_j) in enumerate(val_i)
                if val_j == val_i[idx_j + 1] == val_i[idx_j + 2]
                    return false
                end
                if idx_j + 3 >=  length(val_i)
                    break
                end
            end
        end
        return true
    end

    generate_drops(height, width) = [rand(0:5, width) for _  in 1:height]
    
    ## 指定した欠損条件に対して、モンテカルロ法により、欠損しなかった回数と欠損した回数の組を返す
    function monte_carlo_freq(lack_cond)
        num_ok = 0
        num_ng = 0
        for x in 1:loop_cnt
            drops = generate_drops(height,width)
            if !check_normal_drops(drops)
                continue
            end
            drops = [drop for row in drops for drop in row]
            fire_drops_num = count(i ->(i == 0), drops)
            blue_drops_num = count(i ->(i == 1), drops)
            green_drops_num = count(i ->(i == 2), drops)
            light_drops_num = count(i ->(i == 3), drops)
            black_drops_num = count(i ->(i == 4), drops)
            recovery_drops_num = count(i ->(i == 5), drops)
            if !lack_cond(
                    fire_drops_num,
                    blue_drops_num,
                    green_drops_num,
                    light_drops_num,
                    black_drops_num,
                    recovery_drops_num)
                num_ok += 1
            else
                num_ng += 1
            end
        end
        return num_ok, num_ng
    end            

    ## 試行回数、OKの回数、NGの回数、確率を出力する
    function print_prob(num_ok, num_ng)
        println("試行回数は$(num_ok + num_ng)")
        println("OKの回数は$num_ok")
        println("NGの回数は$num_ng")
        println("確率は$(num_ok/(num_ok + num_ng) * 100)")
    end

    ##=============================================##

    drops = generate_drops(height, width)
    while !check_normal_drops(drops)
        drops = generate_drops(height, width)
    end
    println(drops)

    println("総試行回数は$(loop_cnt)回")

    ## 2色いずれか5個以上ある確率
    println("=========================")
    println("2色いずれか5個以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 4 && d <= 4)
    print_prob(num_ok, num_ng)

    ## 指定2色がある確率
    println("=========================")
    println("指定2色がある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 2 || d <= 2)
    print_prob(num_ok, num_ng)

    ## 6色のうち5色がある確率
    println("=========================")
    println("6色のうち5色がある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> sort([f, b, g, l, d, r])[2] <= 2)
    print_prob(num_ok, num_ng)

    ## 5色のうち4色がある確率
    println("=========================")
    println("5色のうち4色がある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> sort([f, b, g, l, d])[2] <= 2)
    print_prob(num_ok, num_ng)

    ## 指定同色2コンボがある確率
    println("=========================")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 5)
    print_prob(num_ok, num_ng)

    ## 指定同色1コンボがある確率
    println("=========================")
    println("指定同色1コンボ(3個)がある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 2)
    print_prob(num_ok, num_ng)

    ## 指定同色が5個以上がある確率

    println("=========================")
    println("指定同色が5個以上がある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 4)
    print_prob(num_ok, num_ng)

    ## 全色ある確率
    println("=========================")
    println("全色ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> sort([f, b, g, l, d, r])[1] <= 2)
    print_prob(num_ok, num_ng)

    ## ガードブレイクが発動できる確率
    println("=========================")
    println("ガードブレイクが発動できる確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> sort([f, b, g, l, d])[1] <= 2)
    print_prob(num_ok, num_ng)

    ## 指定一色が8個以上ある確率(ねね/覚醒カストル等)
    println("=========================")
    println("指定一色が8個以上ある確率(ねね/覚醒カストル等)")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> f <= 7)
    print_prob(num_ok, num_ng)

    ## 指定一色が4個以上ある確率(進化前リーチェ)
    println("=========================")
    println("指定一色が4個以上ある確率(進化前リーチェ)")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->f <= 3)
    print_prob(num_ok, num_ng)

    ## 7コンボ以上ある確率
    println("=========================")
    println("7コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) -> (div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 6)
    print_prob(num_ok, num_ng)

    ## 8コンボ以上ある確率
    println("=========================")
    println("8コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 7)
    print_prob(num_ok, num_ng)

    ## 9コンボ以上ある確率
    println("=========================")
    println("9コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 8)
    print_prob(num_ok, num_ng)

    ## 10コンボ以上ある確率
    println("=========================")
    println("10コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 9)
    print_prob(num_ok, num_ng)

    ## 11コンボ以上ある確率
    println("=========================")
    println("11コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 10)
    print_prob(num_ok, num_ng)

    ## 12コンボ以上ある確率
    println("=========================")
    println("12コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 11)
    print_prob(num_ok, num_ng)

    ## 13コンボ以上ある確率
    println("=========================")
    println("13コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 12)
    print_prob(num_ok, num_ng)

    ## 14コンボ以上ある確率
    println("=========================")
    println("14コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 3)) + (div(b, 3)) + (div(g, 3)) + (div(l, 3)) + (div(d, 3)) + (div(r, 3)) <= 13)
    print_prob(num_ok, num_ng)

    ## (4個消しで)5コンボ以上ある確率
    println("=========================")
    println("(4個消しで)5コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 4)
    print_prob(num_ok, num_ng)

    ## (4個消しで)6コンボ以上ある確率
    println("=========================")
    println("(4個消しで)6コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 5)
    print_prob(num_ok, num_ng)

    ## (4個消しで)7コンボ以上ある確率
    println("=========================")
    println("(4個消しで)7コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 6)
    print_prob(num_ok, num_ng)

    ## (4個消しで)8コンボ以上ある確率
    println("=========================")
    println("(4個消しで)8コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 7)
    print_prob(num_ok, num_ng)

    ## (4個消しで)9コンボ以上ある確率
    println("=========================")
    println("(4個消しで)9コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 8)
    print_prob(num_ok, num_ng)

    ## (4個消しで)10コンボ以上ある確率
    println("=========================")
    println("(4個消しで)10コンボ以上ある確率")
    num_ok, num_ng = monte_carlo_freq((f, b, g, l, d, r) ->(div(f, 4)) + (div(b, 4)) + (div(g, 4)) + (div(l, 4)) + (div(d, 4)) + (div(r, 4)) <= 9)
    print_prob(num_ok, num_ng)            
 end
            
@time main()

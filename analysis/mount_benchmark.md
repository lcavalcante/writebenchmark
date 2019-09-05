``` r
data <- read_csv(here::here('analysis', 'data', 'mount.csv'), na = c("", "NA", "NaN"))
```

    ## Parsed with column specification:
    ## cols(
    ##   mounted = col_double(),
    ##   id = col_double(),
    ##   time = col_double()
    ## )

``` r
data <- transform(data, mounted = as.factor(mounted))
```

![](mount_benchmark_files/figure-markdown_github/boxplot-1.png)

![](mount_benchmark_files/figure-markdown_github/boxplotnoout-1.png)

    ## Warning: funs() is soft deprecated as of dplyr 0.8.0
    ## please use list() instead
    ## 
    ##   # Before:
    ##   funs(name = f(.))
    ## 
    ##   # After: 
    ##   list(name = ~ f(.))
    ## This warning is displayed once per session.

![](mount_benchmark_files/figure-markdown_github/n-th_latency-1.png)

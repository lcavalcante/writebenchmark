``` r
data <- read_csv(here::here('analysis', 'data', 'task.csv'), na = c("", "NA", "NaN"))
```

    ## Parsed with column specification:
    ## cols(
    ##   container = col_double(),
    ##   id = col_double(),
    ##   time = col_double(),
    ##   mount = col_character()
    ## )

``` r
data <- transform(data, time = as.numeric(time))
```

![](container_task_files/figure-markdown_github/boxplot-1.png)

    ## Warning: funs() is soft deprecated as of dplyr 0.8.0
    ## please use list() instead
    ## 
    ##   # Before:
    ##   funs(name = f(.))
    ## 
    ##   # After: 
    ##   list(name = ~ f(.))
    ## This warning is displayed once per session.

![](container_task_files/figure-markdown_github/n-th_latency-1.png)

![](container_task_files/figure-markdown_github/boxplot-container-1.png)

# hystrix

from Netflix
put it in front of integrations points with other services to protect your service from cascading failures
Handles fallbacks and graceful degradation.
Fail fast and rapid recovery
does it through circuit breakers and bulkheads
real time monitoring and config changes

## how does it do it

* wraps all calls to external systems in HystrixCommand / HystrixObservableCommand objects which execute in a separate thread
* timing-out calls which take longer than thresholds you define
* maintaining a smaller thread pool for each dependency
* measuring successes, failures, timeouts and thread rejections
* tripping circuit breakers
* perform fallback logic
* monitoring metrics in near real time

## how do you use it

you can code it directly or use annotations

@HystrixCommand
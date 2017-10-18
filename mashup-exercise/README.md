# Mashup

## Description

We want to build a mash-up service that takes two web service sources that return an integer, "Source A" and "Source B", fetches the results
from them, adds them together, and then returns the final result.

For example, if "Source A" returns `45`, and "Source B" returns `466`, we want our final web service to return `511`, with the
exact URL of `/result` with JSON structure:

```
$ curl http://127.0.0.1:12000/result
{
   "final_result":511
}
```

## Getting Started

```
$ go run main.go
```

Starts an example "Source A" and "Source B", as well as a hard coded example Mashup service.  

- [Source A](http://127.0.0.1:10000/value)
- [Source B](http://127.0.0.1:11000/other/value)
- [Example Mashup](http://127.0.0.1:12000/result)

Use CURL to examine any of the example outputs:

```
$ curl http://127.0.0.1:10000/value
{
  "Value": 45
}
```

## Problems (and what to look out for)

- *What if "Source A" or "Source B" stop working?*  Make sure your mashup service gracefully handles if one (or both)
  services stop responding.  "Gracefully handles" would be return an appropriate HTTP response code, such as HTTP/500.

- *Bad Data*  Both "Source A" and "Source B" may malfunction -- returning unexpected data.  Make sure that you handle
  data in a format you don't expect.

- *Delays* Perhaps "Source A" returns quickly, but "Source B" never returns.  Make sure that a timeouts are handled correctly.
  (with a reasonable HTTP response code)

## Extra Bonus Ideas

- Can the results for "Source A" and "Source B" be gathered in parallel?

- Given that "Source A" and "Source B" might take a long time to return, how would you implement a cache?  


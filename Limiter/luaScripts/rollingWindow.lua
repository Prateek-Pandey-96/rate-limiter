local key = KEYS[1]
local limit = tonumber(ARGV[1])
local interval = tonumber(ARGV[2])
local current_time = tonumber(ARGV[3])
local uniqueId = tonumber(ARGV[4])

redis.call('ZREMRANGEBYSCORE', key, '-inf', current_time - interval)
local count = redis.call('ZCARD', key)

if tonumber(count) < limit then
    redis.call('ZADD', key, current_time, uniqueId + tonumber(count))
    return 0
else
    return 1
end
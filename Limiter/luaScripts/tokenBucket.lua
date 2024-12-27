local token_key = KEYS[1]
local last_access_key = "last_access:"..KEYS[1]

local capacity = tonumber(ARGV[1])
local current_time = tonumber(ARGV[2])
local refill_rate = tonumber(ARGV[3])

local last_reset = tonumber(redis.call('GET', last_access_key))
if last_reset == nil then
    last_reset = 0
end
local remaining_tokens = tonumber(redis.call('GET', token_key))
if remaining_tokens == nil then 
    remaining_tokens = 0
end

if current_time - last_reset > 1 then
    redis.call('SET', token_key, math.min(capacity + remaining_tokens, capacity + refill_rate))
    redis.call('SET', last_access_key, current_time)
else
    if remaining_tokens <= 0 then
        return 1
    end
end

redis.call('DECR', token_key)
return 0
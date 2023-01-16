if ARGV[4] then
    -- 不知道为什么，这里的split函数，这样写不会报错，要是写到外边（最上边），往下调用，会出错
    local function split(str, reps)
        local r, len = {}, 0
        if str == nil then
            return nil
        end
        string.gsub(str, "[^" .. reps .. "]+", function(w)
            len = len + 1
            table.insert(r, w)
        end)
        return r, len
    end
    local members = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", "+inf", "limit", 0, -1)
    local derivename = tostring(ARGV[4])
    for i, v in pairs(members) do
        local a, l = split(v, ",")
        if l >= 2 then
            if a[2] == derivename then
                redis.call("ZREM", KEYS[1], v)
                redis.call("ZADD", KEYS[1], ARGV[1], ARGV[2])
                return
            end
        end
    end
end
redis.call("ZADD", KEYS[1], ARGV[1], ARGV[2])
local count = redis.call("ZCARD", KEYS[1])
local platformCount = tonumber(ARGV[3])
if count > platformCount then
    local offset = count - platformCount
    local limit = 0
    local delMember = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", "+inf", "limit", limit, offset)
    for i, v in pairs(delMember) do
        redis.call("ZREM", KEYS[1], v)
    end
end

-- redis-cli --eval xxxx.lua

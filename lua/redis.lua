function __split(str, reps)
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
local found = flase
if ARGV[4] then -- derivename 非空
    local members = redis.call("ZRANGEBYSCORE", KEYS[1], "-inf", "+inf", "limit", 0, -1)
    local derivename = tostring(ARGV[4])
    for i, v in pairs(members) do
        a, l = __split(v, ",")
        if l >= 2 then
            if to_string(a[2]) == derivename then
                redis.call("ZREM", KEYS[1], v)
                redis.call("ZADD", KEYS[1], ARGV[1], ARGV[2])
                found = true
                break
            end
        end
    end
end
if (ARGV[4] and found == false) or ARGV[4] == nil then
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
end


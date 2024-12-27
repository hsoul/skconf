local UE = TT
local UF = DD

local Skill = {
    tid = 1,
    name = "测试",
    type = EM.XXXX,
    params = {
        [1] = {1, 2, 3},
        [5] = {4, 5, 6},
    },
    ss = {
        {EM.ETest, {1}},
    }
}

function Skill:XX1(o)
    local a = UE.Abc(0, 100)
    if not a < 50 then
        return false
    elseif a > 50 and a < 70 then
        return true
    elseif a > 6 then
        return ture
    else
        return false
    end
end

function Skill:XX2(o)
    local i = 0
    while i < 10 do
        i = i + 1
    end
    for k, v in pairs(units) do
        UF.DoSomething()

    end
    while UE.XX(0, 100) > 50 do
    end
end

function Skill:XX3(t)
    if UE.XX(t, 0.5) then
        UF.YY(t, 1, 2)
    else
        UF.DD(t, 1, 2)
    end
    return true
end

function Skill:XX4(t)
    if (UE.XX(0.1) or UE.YY(t, 4)) and UE.CC(t, 0.1) then
        UF.DD(t, 1, 3)
    else
        UF.EE(t, 1, 3)
    end
    if a > 5 then
        return fasle
    end
end

function Skill:XX5(t)
    if UE.AA(0.1) then
        UF.BB(1, 3)
    end
    UF.CC(2)
    UF.DD(1, 2)
    local xx = UF.E("xx")
    UF.SS(xx * 0.3)
    UF.YY(3, 2)
end

local State = {
	tid = 1,
	inner_map = {
		id = 2,
		xd = 3,
	},
	type = EM.Test,
	round = 1,
	flag = 0,
}

function State:XX1(unit)
	
end

function State:XX2(unit)
	
end

return {
    skills = {
        [1] = Skill,
    },
    states = {
        [1] = State,
    }
}
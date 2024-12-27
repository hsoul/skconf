skill ack_s {
    tid = 1,
    ds = 1.5,
    df = "测试",
    type = EM.Test,
    ss = {EM.Test, {1}},
    XX1 = func(o) {
        var a = UE.AA(0, 100)
        if not a < 50 {-- sdfasdffff
            return false
        } 
        else if (a > 50 and a < 70) {
            return true
        }
        else if (a > 6) {
            return ture
        }
        else {
            return false
        }
    },
    XX2 = func(o) {
        for var i = 0; i < 10; i = i + 1 {
            UF.DoSomething1()
            UF.XX()
            UF.DoSomething2()
        }

        for k, v = range units {
            UF.DoSomething()
            UF.XX()
        }

        for UE.Do(0, 100) > 50 { -- sdfasdffff
            UF.Do1()
            UF.XX()
            UF.Do1()
        }
    },
    XX3 = func(t) { -- sdfasdf
        if UE.AA(t, 0.5) {
            UF.BB(t, 32, 3)        
        } else {
            UF.BV(t, 2, 3)
        }
        return true
    },
    XX4 = func(t) {
        if (UE.PP(t, 0.08) or UE.XXX(t, 4)) and UE.YYY(t, 0.08) {
            UF.OP(t, 1, 3)
        } else {
            UF.CD(t, 1, 3)
        }
        if a > 5 {
            return fasle
        }
    },
    XX5 = func(a) {
        if UE.XM(t, 0.08) {
            UF.SM(t, 1, 3)
        }

        UF.OI(t, 30)
        UF.AddState(t, 1, 2)

        var x = UF.UY(t, "x_")
        UF.TY(t, x * 0.3)
        UF.RE(t, 1, 2)
    },
}

state sname {
	tid = 1,
	map = {
		id = 2,
		ss = 1,
        [1] = 5,
        4, 6, 8, 6.7,
        ["tsad"] = 45,
        {v = 5},
        {4, 6, 7}, -- sdfasdffff
        [6] = {4, 6, 7},
        testfun = func(p1) {

        }
	},
	tt = 1,
	xs = 1, -- sdfasdffff
	cc = 0, 
    YY1 = func(unit) {
	
    },
    YY2 = func(unit) {
    
    },
}
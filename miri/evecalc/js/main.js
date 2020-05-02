
var lv, exp, fn, dia, tix, wh, wm, of, cosm, pl, st, mfn, mi, kai, ka,
kaif = (a)=>{
    kai = ((b3 % 2 === 0) ? 30 : 60) * a;
    ka = kai / a;
},
gtps, iv, cm, ivsm, mi, gte, ec, ei, ps, lv$, exp$, fn$, mi$, mfn$, pl$, Gte, Ps, Cosm, Ivsm, Cufn, level = [], fnm, cufn, par,
dg = (a)=>document.getElementById(a),
pas = (a)=>parseInt(dg(a).value),
dsa = (a,b)=>dg(a).setAttribute("value", b),
dgt = (a,b)=>dg(a).textContent = b,
b1 = 0, b2 = 0, b3 = 0, b4 = 0,
ole = ()=>{
    if (b1 % 2 === 0) {
        dsa("le", "お仕事");
    } else {
        dsa("le", "ライブ");
    }
    b1++;
},
ost = ()=>{
    let ni;
    if (dg("bi").options[2].selected) {
        ni = 0;
    }
    if (b2 % 2 === 0) {
        dsa("st", "ツアー");
        dg("bi").options[2] = null;
        let tb = document.createElement("option");
        tb.value = 3;
        let tex = document.createTextNode("３倍");
        tb.appendChild(tex);
        dg("bi").appendChild(tb);
    } else {
        dsa("st", "シアター");
        dg("bi").options[2] = null;
        let tb = document.createElement("option");
        tb.value = 4;
        let tex = document.createTextNode("４倍");
        tb.appendChild(tex);
        dg("bi").appendChild(tb);
    }
    if (ni === 0) {
        dg("bi").options[2].selected = true;
    }
    b2++;
},
obc = ()=>{
    if (b3 % 2 === 0) {
        dsa("bc", "２倍");
    } else {
        dsa("bc", "１倍");
    }
    b3++;
},
obo = ()=>{
    if (b4 % 2 === 0) {
        dsa("bo", "自然回復分を含める");
    } else {
        dsa("bo", "自然回復分を含めない");
    }
    b4++;
    //cookie test
    //const man = Cookies.get('pizzamanz_cookie');
    //alert(man);
},
yoso = (a,b,c)=>{
    while (a < b) {
        a += c;
        level.push(a);
    }
};

yoso(0, 58, 2);
yoso(57, 147, 3);
yoso(146, 422, 4);
yoso(421, 581, 5);
yoso(580, 700, 6);

var fil = (a)=>{
    hiki = level.map(function(b) {
        return a - b;
    });
    sa = hiki.filter(function(c) {
        return c >= 0;
    });
    return Math.min.apply({}, sa);
},
sfnm = (a)=>(a < 700) ? (level.indexOf(a - fil(a)) + 61) : 240,
fng = ()=>{
    par = 100 * pas("fn") / sfnm(pas("lv"));
    if (par <= 100) {
        dg("bar2").value = par;
        dg("bar3").value = 0;
        dg("fn").style.textShadow = "0 0 0.5vw #373737";
    } else if (par <= 200) {
        dg("bar2").value = 100;
        dg("bar3").value = par - 100;
        dg("fn").style.textShadow = "0 0 0.5vw #EB2659";
    } else {
        alert("out");
    }
},
nx = (a,b)=>(a - 1) * 100 + 50 - b,
exg = ()=>{
    par = 100 * pas("exp") / nx(pas("lv"), 0);
    if (par < 100) {
        dg("bar1").value = par;
    } else {
        alert("out");
    }
},
kanm = (a)=>{
    if (a === '') {
        return '';
    }
    a = han(a).replace(/,/g, "").trim();
    if (!/^[+|-]?(\d*)(\.\d+)?$/.test(a)) {
        return a;
    }
    var b = Math.round(a);
    return new Intl.NumberFormat().format(b);
},
kesu = (c)=>{
    return c.replace(/,/g, "");
},
han = (c)=>{
    var d = c.replace(/[！-～]/g, function(e) {
        return String.fromCharCode(e.charCodeAt(0) - 0xFEE0);
    });
    return d;
},
rgtp = (a)=>{
    for (i = 0; i < a; i++) {
        let ra = Math.random();
        let rp = (ra > 0.24889867841) ? 48 : (ra > 0.09471365638) ? 84 : 65;
        rgtps += rp;
        //〇□△/回数、〇□70%/回数
    }
    return rgtps;
},
sinzo = (a,b)=>{
    let coltm = ((a + mfn$) < ka) ? kai / ka : Math.floor((a + mfn$) / ka);
    Cosm += coltm;
    pl$ += ((a + mfn$) < ka) ? kai : 0;
    mfn$ = (a + mfn$) % ka;
    let gti = Math.floor(st[0] * coltm) + mi$;
    let gtp = (cm === "l") ? st[1] * coltm : rgtp(coltm);
    rgtps = 0;
    let ivtm = Math.floor(gti / st[2]);
    Ivsm += ivtm;
    mi$ = gti % st[2];
    let gtpp = st[3] * ivtm;
    Gte += (306 * (coltm + ivtm) + b);
    Ps += (gtp + gtpp);
    return [pl$, mfn$, mi$, Gte, Ps, Cosm, Ivsm];
},
cure = ()=>{
    for (; ; ) {
        if (nx(lv$, exp$) > Gte)
            break;
        Gte -= nx(lv$, exp$);
        exp$ = 0;
        cufn += sfnm(lv$ + 1);
        lv$++;
    };
    return cufn;
},
keisan = ()=>{
    //lv$,exp$,fn$,mi$,mfn$,pl$,Gte,Ps,Cosm,Ivsm,Cufn
    lv = pas("lv");
    lv$ = lv;
    exp = pas("exp");
    exp$ = exp;
    fn = pas("fn");
    fn$ = fn,
    tix = pas("tix");
    let hm = dg("tm").value.split(":");
    wh = parseInt(hm[0], 10);
    wm = parseInt(hm[1], 10);
    of = 60 * ((pas("kika") - 1) * 24 - 15 + wh);
    iv = (dg("st").value === "シアター") ? "s" : "t";
    cm = (dg("le").value === "ライブ") ? "l" : "e";
    st = (iv === "s") ? [(ip = (cm === "l") ? 85 * ((b3 % 2 === 0) ? 1 : 2) : 59.5 * ((b3 % 2 === 0) ? 1 : 2)), ip, 180 * pas("bi"), 537 * pas("bi")] : [6 * ((b3 % 2 === 0) ? 1 : 2), 140 * ((b3 % 2 === 0) ? 1 : 2), 20 * pas("bi"), 720 * pas("bi")];
    kaif(10);
    mi = (iv === "s") ? (360 * pas("kika")) : (40 * pas("kika"));
    mi$ = mi;
    mfn = ((wh + wm / 60 < 21) ? 12 * ((pas("kika") - 1) * 24 - 15) + (12 * wh + Math.floor(wm / 5)) : 2087) + ((cm === "l") ? 0 : tix);
    mfn$ = mfn;
    //ec = 306/((b3%2===0)?1:2);
    //ei = 306/pas("bi");
    kabe = pl = gte = ps = cosm = ivsm = cufn = pl$ = Gte = Ps = Cosm = Ivsm = 0;
    let cfs = [],
        so = [],
        soo = [],
        slc = (a,b)=>so.slice(a - 1, a)[0][b];

    //自然回復分計算---
    sinzo(fn$, of);
    for (i = 0; i < 4; i++) {
        cure();
        if (cufn + mfn$ < ((b3 % 2 === 0) ? 30 : 60))
            break;
        sinzo(cufn, 0);
        cfs.push(cufn);
        cufn = 0;
    }
    dgt("ncl", lv$);
    dgt("ncp", Ps);
    //---ここまで
    //リセット---
    let pS = Ps, pL, lV;
    pl$ = pl;
    Gte = gte;
    Ps = ps;
    Cosm = cosm;
    Ivsm = ivsm;
    lv$ = lv;
    mfn$ = mfn;
    mi$ = mi;
    exp$ = exp;
    //---ここまで

    if (pS === pas("gl")) {
        dgt("rsf", fn + mfn + cfs.reduce((a,b)=>a += b) + "(=自然回復分とレベルアップの分)")
    } else if (pS < pas("gl")) {
        //自然回復分より大きい---
        sinzo(fn$, of);
        while (Ps < pas("gl")) {
            cure();
            let iti = sinzo(cufn, 0);
            cufn = 0;
            soo.push(pl$);
            so.push(lv$, iti);
        }
        let io = soo.indexOf(pl$) * 2;
        pS = Ps;
        pL = pl$;
        lV = lv$;
        lv$ = so.slice(io - 2, io - 1)[0];
        pl$ = slc(io, 0);
        mfn$ = slc(io, 1);
        mi$ = slc(io, 2);
        Gte = slc(io, 3);
        Ps = slc(io, 4);
        Cosm = slc(io, 5);
        Ivsm = slc(io, 6);
        kaif(1);
        while (Ps <= pas("gl")) {
            cure();
            ni = sinzo(cufn, 0);
            cufn = 0;
        }
        if (pl$ === pL) {
            pl$ = pL;
            Ps = pS;
            lv$ = lV;
        }
        //---ここまで
    }
    dgt("rsf", pl$);
    dgt("rsp", Ps);
    dgt("rsl", lv$);
    dgt("cos", Cosm);
    dgt("ivs", Ivsm);
    dgt("rmf", mfn$);
    dgt("rmi", mi$);


    //cookie test
    //const pizza = "ピザマンください";
    //const anko = "アンマンください";
    //Cookies.set('pizzamanz_cookie', pizza, { expires: 16 });//https://github.com/js-cookie/js-cookie
    //Cookies.set('anmanz_cookie', anko, { expires: 16 });

}

//レベル変更時、経験値と元気のゲージを表示しなおす
dg("lv").addEventListener("change", exg, false);
dg("lv").addEventListener("change", fng, false);

//dg("lv").addEventListener("change",()=>{dgt("nx",nx(pas("lv"),0))},false);
//dg("lv").addEventListener("change",()=>{dgt("fnm",sfnm(pas("lv")))},false);

//経験値、元気を変更したときもゲージを表示しなおす
dg("exp").onchange = exg;
dg("fn").onchange = fng;

//読み込み時ダイヤの数にコンマを付ける
window.addEventListener("load", ()=>{
    dg("dia").value = kanm(dg("dia").value)
}
, false);

//ダイヤ入力するところをフォーカスしてないときもコンマを付ける
dg("dia").addEventListener("blur", function() {
    this.value = kanm(this.value)
}, false);

//フォーカスしてる時はコンマはずす
dg("dia").addEventListener("focus", function() {
    this.value = kesu(this.value)
}, false);

//ライブか仕事か、ボタン押したときに文字を変える
dg("le").onclick = ole;

//シアターかツアーか、ボタン押したときに文字を変える
dg("st").onclick = ost;

//元気消費倍率をボタン押したときに変える
dg("bc").onclick = obc;

//自然回復分について、ボタン押した時に変える
dg("bo").onclick = obo;

//計算実行
dg("ksan").onclick = keisan;

/*post---
function senpo(){
  let xmlHttpRequest = new xmlHttpRequest();
  let formData = new FormData(document.forms.xhr);
  xmlHttpRequest.open("POST", "/test", true);
  xmlHttpRequest.send(formData);
}
//---ここまで*/

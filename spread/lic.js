function s(t, n, e, i) {
    if (t.length) {
        var element = t[e];
        var i1 = i(element);
        var ts = t.splice(n, 1, i1);
        var t1 = ts[0];
        var i2 = i(t1);
        t.splice(e, 1, i2)
    }
}

function d(code) {
    return String.fromCharCode(code)
}

function u(t, n) {
    var n, e, n = 1 < arguments.length && void 0 !== n ? n : 1, e = t.charCodeAt(0);
    if (65 <= e && e <= 90) {
        return t.toLowerCase()
    } else if (97 <= e && e <= 122) {
        return t.toUpperCase()
    } else if (48 <= e && e <= 57) {
        return d(48 + (e - 48 + 10 + n) % 10)
    } else {
        return t
    }
    // return 65 <= e && e <= 90 ? t.toLowerCase() : 97 <= e && e <= 122 ? t.toUpperCase() : 48 <= e && e <= 57 ? d(48 + (e - 48 + 10 + n) % 10) : t
}

function c(t) {
    for (var n, e, i, n = t.split(""), e = function t(n) {
        return u(n, -1)
    }, i = n.length - 5; 0 <= i; i--) {
        s(n, i + 1, i + 3, e)
        s(n, i, i + 2, e)
    }
    return n.join("")
}

function f(t) {
    return t.split("").reverse().join("")
}

var g = function (t) {
    return Buffer.from(t, 'base64').toString('binary');
}

function m(t) {
    return g(t)
}

function y(t) {
    var res = m(t)
    res = Array.prototype.map.call(res, function (t) {
        var dd = "%" + ("00" + t.charCodeAt(0).toString(16)).slice(-2)
        return dd
    }).join("")
    res = decodeURIComponent(res)
    return res
}

// XRsZ => Evl
// HUkJ => Prd
// T&g => N
// Q&w => C
// GRz1 => Dms
// JYx3Gb#8Pb5R => location
// VdgJHc#wJb59 => protocol
// 4LJITMx8UMcA => 127.0.0.1
function M(t) {
    var n;
    if (t) {
        t = c(t)
        t = f(t);
        n = Math.ceil(t.length / 2)
        t = t.substr(n) + t.substr(0, n).replace("#", "=").replace("&", "==")
    }
    return y(t)
}

var u1 = ["XRsZ", "HUkJ", "T&g", "Q&w", "GRz1", "JYx3Gb#8Pb5R", "VdgJHc#wJb59", "4LJITMx8UMcA"]
var u2 = ["Evl", "Prd", "N", "C", "Dms", "location", "protocol", "127.0.0.1"]
for (var i = 0; i < u1.length; i++) {
    // if (i > 0) break
    var s1 = M(u1[i]);
    console.log(u1[i] + " => " + s1 + " " + String(s1 === u2[i]))
}

var licStr = "2c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"
var lic = M(licStr)
// lic = JSON.parse(lic);
console.log(lic)


function A(t) {
    var n, e, i, r, o, a;
    n = 0
    e = 5381
    i = 0
    for (r = t.length - 1; 0 <= r; r--) {
        o = t.charCodeAt(r)
        e = o + ((e << 5) + e)
        n = o + (n << 6) + (n << 16) - n
        i = o + ((i << 5) - i)
        i &= i
        a = n ^ e ^ i

    }
    if (a < 0) {
        a = ~a
    }
    res = a.toString(16).toUpperCase()
    return [a, res]
}

// 615274881    24AC5981
//
var vl1 = 'E879948536774266#B1{"Anl":{"dsr":false,"flg":["ReportSheet","DataChart"]},"Id":"879948536774266","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"127.0.0.1","Exp":"20250606","Crt":"20250507 032315","Prd":[{"N":"Spread JS v.18","C":"BJIH"}]}'
idx = vl1.indexOf("徽")
var vl2 = vl1.substring(0, idx)
var res = A(vl1)
console.log(res)

// 签名密钥
var pK = "l6/zrbWoSbcLFwEetFh38rH3ErBZE9H+Cqix3R+wTlfA1wD5B+lUcCQn+EJ60I4RGrm0x1sFjkiLWwB0jAn6BWZv0W4WbqAKriOdeoivxDp1Wmjs3qkEDhvbsjPtfvwx2BHil6o+/tDrdMJQSGs18WZm2PoQLQuL+9VhZ4FNRHUQU3Jtioke/OZEGHJOdYVwvCGalzBad6QFOiVbDBQPePpS3++GJzOxN8SN/7lyS5/IdKiy3WJRaVGkB370+HbN6hKraDfUgReLX26yxRaKC/5aWnGAJ2NnWLoGyAGRcwT9dVjo4bcAZNrrA0U9JVKQxaSskhdv2p49XzJkltXx5w=="
function x(t) {
    return t.replace(/\+/g, "-").replace(/\//g, "_").replace(/\=+$/, "")
}
var  n = x(pK)
var sn = {
    "alg": "PS256",
    "e": "AQAB",
    "kty": "RSA",
    "n": n
}

console.log(sn)
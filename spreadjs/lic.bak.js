function s(t, n, e, i) {
    1 != t.length && t.splice(e, 1, i(t.splice(n, 1, i(t[e]))[0]))
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
    return decodeURIComponent(Array.prototype.map.call(m(t), function (t) {
        return "%" + ("00" + t.charCodeAt(0).toString(16)).slice(-2)
    }).join(""))
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
    return t ? (t = f(t = c(t)), n = Math.ceil(t.length / 2), y(t = (t = (t = t.substr(n) + t.substr(0, n)).replace("#", "=")).replace("&", "=="))) : ""
}

var u1 = ["XRsZ", "HUkJ", "T&g", "Q&w", "GRz1", "JYx3Gb#8Pb5R", "VdgJHc#wJb59", "4LJITMx8UMcA"]
var u2 = ["Evl", "Prd", "N", "C", "Dms", "location", "protocol", "127.0.0.1"]
for (var i = 0; i < u1.length; i++) {
    var s1 = M(u1[i]);
    console.log(u1[i] + " => " + s1 + " " + String(s1 === u2[i]))
}

var licStr = "2c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"
var lic = JSON.parse(M(licStr));
console.log(lic)



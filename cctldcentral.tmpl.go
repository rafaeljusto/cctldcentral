package main

var ccTLDTemplate = `
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="author" content="Rafael Dantas Justo">
    <title>ccTLD Central</title>

    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/2.8.0/css/flag-icon.css" >

    <style>
      body {
        padding-top: 50px;
      }

      .sub-header {
        padding-bottom: 10px;
        border-bottom: 1px solid #eee;
      }

      .navbar > .container-fluid {
        background: linear-gradient(to right, rgba(30,87,153,1) 25%,rgba(255,255,255,1) 75%);
      }

      .navbar-fixed-top {
        border: 0;
      }

      .navbar-inverse .navbar-brand {
        color: #fff;
        cursor: default;
        font-size: 20px;
        font-weight: bold;
        text-shadow: 2px 1px 4px rgba(170, 175, 255, 0.75);
      }

      .sidebar {
        display: none;
      }

      @media (min-width: 768px) {
        .sidebar {
          position: fixed;
          top: 51px;
          bottom: 0;
          left: 0;
          z-index: 1000;
          display: block;
          padding: 20px;
          overflow-x: hidden;
          overflow-y: auto; /* Scrollable contents if viewport is shorter than content. */
          background-color: #f5f5f5;
          border-right: 1px solid #eee;
        }
      }

      .nav-sidebar {
        margin-right: -21px; /* 20px padding + 1px border */
        margin-bottom: 20px;
        margin-left: -20px;
      }

      .nav-sidebar > li > a {
        padding-right: 20px;
        padding-left: 20px;
      }

      .nav-sidebar > .active > a,
      .nav-sidebar > .active > a:hover,
      .nav-sidebar > .active > a:focus {
        color: #fff;
        background-color: #428bca;
      }

      .main {
        padding: 20px;
      }

      @media (min-width: 768px) {
        .main {
          padding-right: 40px;
          padding-left: 40px;
        }
      }

      .main .page-header {
        margin-top: 0;
      }
    </style>
  </head>
  <body>
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container-fluid">
        <div class="navbar-header">
          <span class="navbar-brand">ccTLD Central</span>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
            <li>
              <img height="50px" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAMgAAABkCAIAAABM5OhcAAAgAElEQVR42u1dd1xU17Ze55ypwNB7B0WsEcUSLKHojYqKXWONsdxoknt9NyYvzxc1mlwTn7k3mmhiF1sUEzWxk1iiXERRFBAFBAUZBhgYYJh2ppyy3x/HEESEmRGMJuf78cdwzu7n22uvtXbDEELAg0d7A+ebgAdPLB48sf5YYHh9gSdWRyD2vGLMFWW5keGbwkpgvPJuDUwMcj1Y3M3H4UG18bVe7ht6uYtxjG8WnljtgDIjE5cqjwlwPFOqa1AZO3lKegc5vdXNbZi7mG8cnlhPhc2lupT7mvX9vUUC/FSZLvluw10lKRQR4zs7T4p0nebrwDcRTyw7MSC1/HqZTjE3MkBMAMCWEu0PKqOLgf7+ngZwfHSYbFqEy1h/R1dBGwPlsQpDhYl5q5MzTyweDyHeWdjTW3pjbMicq9X7MmsAQVSES/bo4At15jVZNRnlepOBjghxSgySTQxydHcQBEkIF8EjRtIDkp6cpkzqJBvqLo73kPDE4gEAcLSSnPT9/Q+GBW69Xa9TmwK9HcqqDD3DnW8lBnNiKktj+bJAfVFJqmtNRpplEYAIx2WiCCeRBIO7FsakNC4e5LO2q6uz8A9rlfPEsgcTL1X9kF8/vJfHuZzatOmdE89X6KuNq+L9P+rp3iwkjVCVidXTLACIcCxYSrAAPkdK1FM6/bGtSp5Y9oBGIEkuZCgWKOTgJtLPjJDtLDQYqIuzusS2ZSSuLFCrabSxl/sfu4l4B6k9EGCw+9UgMNIAiGwwZ6nN6VPCgcAmn1W0Gff0fe3/dnb5wzcRTyw7MSvQcURvTwAABHEnyno7i2b0966tNLybU9d6xDwj4yHGeWLxeCIODfEFMQEApMbybbn+2wHewUFO669WV1JPnPkxMIjWWgQYxhOLxxPhIsSzpoQDAGAwO7WcZNHxEUHAoPE/VzwpisrE4NI/A694Yj0dol3F6/4SBAjARO++r+0tE36W4H/9vub7KkOL4b8r0/X3kdrEKxZBiZHW0r/ZWJ/mqwHgVLXxB4WBfl5NL94qbAdMu1n73RUlyERoThcA6HyyTKkk9Qu6PR5yUmaN0kBdTgiwPvFqE+O7NR8kxPejQxgWlZG0m0xYVG/+14UKILCDSaFRriIaQTdnEfE8CUIBT4t2ULb6ev7FXZzXYOb+vToiyGtnYY+zijt/CWwWsoqkDRTbYiIMAgYhId58nJQSGCbAkIWVifC8OtPyK9U+bmJlFdm7q+uPQ/2mZdVeu10HNHt+dpcENzFPrD8aFoTKAGTcb08hfnNap777i//mLd3Y26NpsGCZ8NgdPQJoyp47OmpselVFJWkx0QI3if61Tk3X5GAASIC/OdCnj4tIhGM7hwcO95Z6iXFOUysobgAWAYD6CXzldaw/FPq4iefH+Gy6WOm0q/CtnNrG5/MCHSmSatQ+ik1M2OGSnnvuKjWW46OD6xZ2xwyUkX1EOXES4g3zu26J8vAWE/GektnBTn4SolH//0tXV6ARSAVDXUW8jvVnweVa05zLypISrcBRkNDN7dsB3u5igthV6IJjk7u5yUnqbEEDIRWcGhU0wtcBAD7IVz/QWg697GNTLrlaS6STUPKcLTzkidXhSKszbbit/uFuA9DsuC4uceHOWwoaipSkENDHQ/zei3QlAACAQki06+7clzx2Z6teCZZdGhn0QteaJ9YzAoXQbgV5slCdKtdbcEzqIKBpNsxBMLeL66QwWRcHAQAYGRRzumxWV7dlFyoSI12PxfnzxOJhA+4b6CNy3XG54b6KVGoosDAhobJP+nhODHJyJDAA2KswvP79/aw3ukY/Z5oTT6wXBhYW/bOoYWueuqaGBAxbGOPzSTc3HwnR48dSDMduJ4XyxOLxVCg10MMuVJSV61kWxfZwXxAum32qTPvX7jJBy5Y7g6DEQOlpFOUqepLeTrOQLNctDJXxxPqz447GHH+1RnVfiwFCNNyY26Wvcwuj4dV6U9I1VaBEMD3Q8f3ODxfOswhh2CP+1bjU8kv3tYVzu0Q6Cnli/alBI/imVPduVg1jYcFAiQjs7YE+XzRZFfiTyviP7LpqAnsnxGlRoJOJYcMchQhB3NXqBA9Jeo3xllz3z1j/tVdryg0UpadAgK+L938/woUnFg9gEFyqNX6YW3c1Xw0ADg6C1XH+Xg7C5WmVChpe6+6ysYe7p5igWPiyRGNm0Kk68+VBPpysemCk/1NjipQJpUL8/es1GhNzJTGYHwp5PILzKuP0C5UqlRFoBI6CHQkBr/hKI5qMa9UmpsJE93YWES05SC0AfY6V+UuJs68G8sTi0QIWZql2ZNb8OCF0nL+jrXFnXak+VabPHBvSxVHAE4tHc8y7VZ+SU0vO6WJH3KPVxkUXKuZ3c6NZtoxkYr0kg/0copxFPLF4wCWVKe5YacWcSH+RnUsH/qM2by7SXK0wKGtNRpJ28pX+8mpQv445e4JfNvPCIF1jBhbcBPZPNg91Ew8d6M393laqe/NCRf9vi+cN8tkZ7dXupeWXzbwY+L+ihuVpyjVDfKXttIrhr2EyZl7XfpGuu65Uf1LYwA+Ff1KEHimNDXTc86u8aS/ISTrk0H2xg8A0rRMvsf50+PKBLsJV1L6symqwzLlVPzFDCRbGjOBwFflH0bFMZqirA70BcBxkMvD15gn0JGy7o/52kE/7pnm+1nRdrnND8OnwgHPF2hWFajcJ4SMieraTS8L2ofCTz+B2PiRvBQe7zhnTauHQYTh0GEofQG0dUBQAgFQKvr4wIBpmTYdh8XZWxWCArJtQUAglpVBVDWbTw+denrBqBby9xOYEaRreehOGJwDLwjv/gNpaOLAHBNa1+5EfIeX7R1a2IwBHKXh7Q0Rn6BIB/aOtb8AvS3QihBa362FaGWrzj5WkswhPrSIv59UDQkCxIMA7BznO6eQyNNBhqKvkabb92E7POwWQngGMXce87jsAS/8HGAZcXSEsDMYmgkQCgEFVFRTdgzM/w9HjEPUSJG+D8FAbks28Dlu2w7ETwCAQCUEiARdn4GZjEUCgDigLFBbB410Iw6C2DhgGvFsyi2gG6tUPf1+/AVVKYK3esFCugP9cBokYsCbKBmKBNALLAs0ADjBlEryzGHr1aDOxJeGysOMP2pdYg9zEg9zEALAkTOacVwcIgMAA4F65YaVcDwwCF9GELq7vRDgP8JQ62U4x24ll3zZehGDZSti1BxwcYOkSeGdRC+nU1sGb78DlKzA4Hk4cgX59rfp+C96CrBsgFkOXLjA+CV6bDMFBgD+mO9663nIK0YOhqAhuXgFhK/P/9to3J45Arx6/xUYIDAZIvwJHj0F6Bvx4Ag7/AC8PgO3fgL9f6ykZARoo1rUDztOSCXDlgm77q40lanNVvTmPpIxm1mxkzGb65wL18Vt1DGADu7m8GeIc5SYKcRK6W1kGZCumz0GBnZFGa1uszduRbwjqHtV2xB9PIO8gFBzRdsj9B5BPMPILRa+ORXIFsg/9BiFnb2Q0thaGYdCAISioMzKbrU32y03ILxTl5rUW5twF1DMa+YUiv1D03ZHW0+t7sSJHbULPHCTNPjBQV+tNK+82QHIhbMx77Uq1NRFt7wHYo5virEGZHFZ9DGIRZGWAc1uLzsaNga2bwGyCpInQiv63bAW8vwwAwY5v4KfjEBRgZ4fFADDU/rW2JvyweMjLgg8/AJqC/1oKK1a1ErZGabzZYHn2doOUwEIcBAPdxKu7uKC5kb1CZCk3aoxW3KeA2/UpbMTuvYAT8NUXIJVaFX58EvSJgnslUFPTcoDN22DvfhCJIf0XSBz1lOaLdSywoz9ZF/7tRZB6AlgWdu2FzVtbDFJlpBO7u24v0vzu9imGABAA6ghiYQA4ZkMjsyzs2Q8e7pA0xoZcPv4IAMHGb1rWqz5bBzQNZ45B56d262EYWOPLxq0L1jRZ66nYtw+cPQ0MDZ98BjezH3/vJxUUmdgrJdpbOup3ZFUpSd8q01qpZ3e8xLpfAgYDdAq3LVZkBDAMpKU3N8RYFqbPAZaFDV9At67t0Qc7qm/bZuX07AEb1wNBwMzXwWh8/P1CfwcQE/99reZ3JNbWIg0A5ugqsuZWDvuIZUssRSUgDLp3s3FsdwCxFBRVDx1djcjOhdIy8A+AqRPbjQFAWBEMt72tcNtoO2USREaC1gAHDj3+ckaAY0SEy093G1JrjL8Lq9JUpv9LrwKADYP9rBLxHc0rqKsFHCA8zMaPgoGLE5iMzR1m//oCBDh8/ikQRLu1GY7av9YPZaGNfoo9O0BIwJebgKZb0FS7u4GHZN65imfPKjMLo1PlgGDuAO8FYTIre5WtnwFDNgl5BMg+1xeGoWZdQ6OBtHQkEsHgQe3HKgwwvA0GIADM5log3HZ9P8AfBQagGhVk57Tg0nQX/29P96o604cdsBihdbx6QaFXm6NCZclWz1faTCwEYOt5Ofadr9NCRvdKGAyDyAgQtJu4Yq0r3jOrNTYuiRUI4GJai2/X9HBL6O3x6aXKciP9zFh1VEmm5aslMlH22BCb9IDnRt9tE7m3AMOwoUPgD4yxiZiAQBlXn/T+/Ct+c/t6ddpfLCefEbeW/FIBCA4n2nZIic1TOgizTcg/DM8FNxqtVTsY9vGMUHk5InAY2L8dWw3hGBBteU8wABxHNrobkH0r8rpGsiIhVl7eikxO7u/lJyZCkgvHR3keGeTboecX5Wgtinrz+Jc8Rvs4dDCxbJ85Q9y3YVk6qj+YjNZGIR4rW0UlwgBCQtq37VjUtlnI2ljrX1vJrklGBwcwmYBlW5jx/BWfvuT+erhsdrqS2J6fPz2im6yjNjqfqTaBmX2rh5utEe1YfIOhNrXdx8MjDDCM+O+lwFineyDEfLWJ1T9y9jBiGQDc2oUr1jKAU97b3/OO7LEkOfWEQMC06d2OdBJeGxm0+4Gu+6F7qtcjPTvmvicXMQ44VNt+Z7EdQyFmWwvjnBsaAwzD5r9hdTYI274DDIameWF+fggDUCggqP22X+KYNfOliADbhjacc9bbxXWTEZNIrPSnzA2V7QhyWn6zdsvADlkp+Vaw07tOwv33NLNCnMDGz26Pvdbh9hHLIq7To0etcQC4dbtdJZa1xbNHAbBjJKyvRzQNnh7Wx/j8ZZ+t+fUdNBSeUJJmIzPA0+ZLFW0mFgsYa+OgwNplR7KANf/kUVEMAvrc+fYdCpEVxUM21xqzr9Zw4SJD0VjPHtbHiHERiQlcYWLal1JmFn1+T5N0+H6Qv+Oq3h62RrddX8FtnquwOfwTMsLCQzEMMSUlAoZpN887l0ubosVWfQmzs9bMufMIQ1jfPjbF6uQsuqajAiXt0yZ6Gu2W6/6WrsQMVHR39+sJAXZ8PUEHdfHHwmPtkJG3t6BTZ0ouh7vF0L1ruzSilQLY1lr/qrzbSivGci0LWSg8LtameBICM+os4PW0twAXG6jR/1He01OIRl19Hc4P9fWXCuzusHZZUh3dxE/4loIP3kMssrz/P9BuuyGtHQqRbQsWMDtqjXYmszq9cFAMBNq2bjFbSYLgaa3CAwp9XGr5gnDnf0d71UwKK3g10G5W2SWxhATLWMBsASutBAyQABCGbG1mhAN6TLRjrwzBZI6We0WirCzo3689hkIERFvyFMNAKkY4AqOx1aXxzQ1J20DT5NZtLKKF7//DpnjDzikQSQc8tStr5tkK+fTOviJc2B4uV5tpLugcgVgE53+xUePA2kVigUjksPUbhOGGvy6GhnaYi7VqKMQwvEsXxLBQ+sD6ZG1VACz/eI8ljeKBL2P9bZhaMDLoQpEGAPJVpqdphx0P9Kv7ewdJCGE7OfJtJhbx8gCEwLh3v/VW91Ns4m+pkv37SSZNZEmjcfJroHn61bpWtaNg0gRAYPn3hg6y6pk9+8w/ncUIQrTh3zZF3K4wAIMAA83T3aWz6YF2ZU+3dqyRzcTChgwmAv0tD0rR1avWOooIHNneDVgMY5/Qe4RrPhYmxFsqKgxjxqGSkqeSWBggvO1GwKL7IpHAeP0aKr5nJV1Z3Nrhn966XffJGlYsdDqwH7w8bSp/rpLklNi79q5aZhGMulaTMdSvfbuKPRqf9LM1iGa0CxahKmXbLYxh9gkshLXmAJNs3igZO5qqq9ckJtEbNoLeYK/Awqzq6VKp09pPEU7oXpsJqtp2M1mU1eT02br1X2IymcvOHfBST1uLf0ZhAAwDBI4ie4awGjPT6azi9ABvh/a+7NAeYuGDYpze/CuLoCHhVWbft0BRrYx2Dx3odulYrQcQr1vr/OV6zMFRt32neuBgasUqKLoHNA0sa/3oa70fAU8cJR05gjGa1HHD2ZOnAaFWa90kWS5k4x/DAEWhq5nG2W+o44abb+UJQkJcLpyFAfbYIiR3SSKOvR3p2vq48ThSyvUvn60ofTWwI5ZH2GlPCv7xd+dOYdrlH2nW/Qtb/yXh7o5LHjkYDpnNsh8Og5MTAIZwu4QW3vbKE/wvw1zi0pkjP+g3fq07+gMcOw4EQbi44FIJ9qsHFff2dNi540meAmTL+hbxv9cR3bpqN3zVsOxD7KPVhLsbLn6s1mdOgYAA/Ne5xdpa7dQZv+WOASJJxkACywLDCH29nf75CTY4xr795Uaa1VoYAOga5NS91ZPcd8v138p1e6K9AqQCACgm6V4nyr7q71UypqMOVLbfUUEkjXWLj6NTfzZ9e4BWKpFe36yJH3ZoAiecHO1pOLEYd7LCpSEUEq9NdZk0AcnlTE4ulXrWcreQVTf82lMR0epqCEwsIhwdrPeKCRbMc580kf7xR9PRHxlVLaNrVmvTw+2vOE44OnKLhRBpeERgCwXSgQOEo0YS/aMxP7+nmUIoJWnEIKDZ1a0eyYcARADnc+s/dBJ1dhL++6ZqTJisemKYi7ADD7HiD157gbHoVt3WS1VevtKaKW3sr/yu0jDthHxAqNP4IKeFnV08RR1+Lhp/BukLjDOlOkBo75C2DbqJfg4+nuI9MT4FJPMMWGWP8q5SqcrLy9sMlpeXZzK15rIzGo03btywr9A3btxg2Y66Apmm6dLS0mYP6+vr9+/fr1KpWo97+/Zts9lMkqRcLm98aDAYKisr29/1hUDZYAapIMG37YMLBBj2fWzA7CzVkUrDsyG9zcT66aefvvnmmye9/eyzz3Q6HQAcP35c06r3Mjk5OT7enjPWEEIpKSkMw3RQi+j1+hUrVjR9IpfLo6KiAKB79+5KZWseltOnTxsMhtu3b3/++eeND/Pz87dt29YBxEKMgf5rtKfIOv3VW4JnKclgKWFm0PNILIIghE3my1iW3bVr19atWy0Wi16vz8jIOHnypMlkGjZsmEwm02q1aWlpa9euvXv3brN0Ll269Omnn549exYAtFrtxYsXV69enZ2dvWPHjj179nBhsrOzV65cef/+fe7r3r59e//+/RiGjRkzhiAIAEhNTV2zZg3HYIvFsmXLll27dlEUBQAZGRkpKSl79+59XJSuWrXqzp07AKBWqzMyMlavXl1RUQEAFEV9/fXXZ86ckUgeWSawbt268+fPz5o169q1a9nZ2QBQXl6+evXqtLQ0Lt/i4uK9e/fSNB0XF+fg4CAQCLRa7erVqznJh+N4fX39xx9/XFRUxCWYkZHx0UcfKRQKrp8cOXJk3bp1RqMRAK5evXr8+PH169fTdBubcPI0FoZFSyNa8zKwAMU66r9y64KOP+h6oPjAK37lRkZMPJNtVrYemHTgwIEVK1Y0/pucnHzixIn9+/e/9957DMPMmzcvJyeHpumYmJiioqK8vLyePXtWVVUFBwc3S2fQoEGlpaXjx49HCGVlZY0ePdpgMHh7eyuVysWLF1dXV5MkGRERoVarubi7d++WSqX3799nWTY4ONhkMv3yyy9vvPFGYWHh2LFjEUJ79+7NzMxMSUnZuHEjQigsLKygoGDevHk1NTWNmdI0PX369IaGhs6dOyOEMjMzp0+fXlFRMXXqVITQ+++/f+jQoZ9//nnhwoVNixoXF9es8K+//rpGo5k6daper6+pqREKhZcvX+ZClpeX5+bmRkRE1NbWurq66vX6Gzdu9OvXT6vVcpnW1dXFxsbW19dHRkYihNLT09etW3fr1q3Y2FiE0JAhQw4fPrxmzZq9e/e2/iH2yXWw+Y6BZlt8qzYzK/LriX1FsC3/77dqM+pMmfUm3wsKM8M+m4O1nlZ5nzNnzoQJE7KyssaNG4fjuIuLS3BwMEEQOI5jGEZR1LBhw5ydnZvZnhcuXCgvL//4448vXrzIPeH6ekREhI+PT0hICEVRCoWCoqgZM2YoFIr6+nqKopYtWxYeHo4Q4hLftWvX3//+98jIyGPHjgHA1KlTExMT5XL5/PnzAcDf379r165OTk5ms7mpuB03blzPnj25hyzLzpo1y93dXa1Wc4Jk3bp1DQ0Nhw4datb3mtV68eLFMTExarWak44zZ84cNGgQN83ASb6xY8d6eHgMHjy4qqoKAMaPHy+TycaMGVNTU1NaWlpTUzNr1qwHDx6QJDl48OB9+/aNGzfOy8sLAEwmU3x8vMViaWhriv2I0ggyYYse8zdu1u2+WuXoKU19NdBAo30K/U09NT/AsTI+4JntCbVrw2qTQT0hIWHNmjVHjx4lSRvOc96zZ09ycvIXX3wxbtw47qO25J8Sjh49+ujRo/X19e7u7gghFxeXph/bzc2trq4OAAoKCgAgMTFx7969Bw8ebMqkx5XrU6dOlZeXBwS0sODJ0dGRI18zJsXFxWVmZnLj1MGDBwHggw8+uHPnzqJFizgbojG1xpbhCEeSpEgkaqSmXq8nCEIkEs2fP//w4cMqlcrBwWHVqlVJSUl5eXkUZdtkX2qZLjaohQubzCzanVHl5uNQPiFsqKdklK/08ADvt4OdJAT2LHca2yyxBALBwYMHc3NzcRz/7rvvoqOj33vvPZVKNWDAAAAICgr629/+9vXXX3M6EIZhOI5zekbTRIqKimJjYwUCwcyZM3ft2hUfH899kqaBg4ODT506RZJkdnZ2ZmZmUzYLBAIAWLlyZXx8/JYtWzw9Pbdv3x4dHb148eKGhoaRI0c2S6oR/v7+aWlpo0aN4hj5W/fCcQBYtGhRnz59hEJh376PHH+6ZMmSAQMGxMfHnzt3jtOxdDrd+PHjr127tmjRIgz77VpTrtYikWj37t1qtdpgMAQGBtbW1n7++ed5eXmVlZUeHh6c6MrPzy8sLExPT+/du/eKFStwHOcq1VhgrC0amLSW2f1b2JkjxrGR/bxTM6oqjHRP59/vrnJbx06z2Vz3K7gneXl5FotFq9UihBiGycrKomlap9MxDEPTNEmSLMtqNJqmiXCBufA6nY6maaPR2PjcaDQyDMOpRNevX6dpGiFkMplMpoeHcGo0GpZlEUINDQ23bt1qTDY3NxchpNPpGpMiSZJLqhEqlUoul5MkiRCiKMpisbAsy0VBCJWUlKhUKoPB8HjFuXpxvw0GQ0FBAcuyXHQuNS5rrtYmk6mxYBRFmc3m3NzcxuhmszkrK4urAkLo7t27BoOhoaGBKzbLsmazubGyLUJLsbD5Tr6mhTNR52QoYUNuilyHflfwnvcXEj8rjYlHSsjF3UWPTnSOSqtKvV1/emLYKF+H37eEz++VJ23a239mLMurc/eWCh5l1X9qTanZtf0jXX0lv/+ESnsSS6lUjh49etSoUUlJSYmJiQDw1ltvNXoyaZoeOXLk+PHj582bp2lr5efatWv79u07c+ZMi8XOo4IRQosWLWozzJUrVzhbbOXKlS8Kq+7oqJv56qwxwc0+npykMAIrkus9RM+BvLBj+GQYxmKxNP5LURRFUY3/fvHFF43qRb9+/Rpfmc3mTp06cVpO9+7dGx1LTZPicPbsWc6TxHlNuRxZluWS4qJQFMXpKE0LwzBMs7IhhLiINE1zyhanwTTqcD179myWO6d4NSbYYgl/X4y9WPnymbLHnw88q+h0suw5KaTN1CZJMjw8fPLkyatXr+Yc4r169erTp8+ZM2caJVOjlHpcgWNZ1tPTs0ePHnfv3jUajf7+/rNmzXrnnXeahjl27BiXeGxsLOc6f/vtt0eOHLl8+XK9Xh8YGLhgwYLQ0FClUkmSZLdu3SZOnMj5rpYsWRISEvLKK68cOXIEAGJiYgCgtLQ0Li4uJCQkOTkZACZMmDBlypQePXpwXoN79+4tXbrUYrFMmzYNAHJycrjqcFMCkydPjoqKiouLy8vLe34k1tly/bdD/R9/nllpGOwrfV5KaSsTv/rqq82bN7Mse+PGDYRQTEzMvXv39Hp9aWkpF2DdunXZ2dnc7+jo6GYSixMV7777blZWllwuDw4OpihKr9c3zWL06NGNTwIDA1mWnT9//oYNGzhhdvz4cYRQ7969OVGkVqv37dsXFBSEEJo7d+6NGzdycnI++eQThNBLL73UmGZ4eLjZbOYMyZs3b3KyszGMXq9PSkpCCCUkJHDmWExMDEIoNja2oqLi5MmTKSkpz4/ECjtSMuNaC9dDwNe3J2QoX1SJpdFounbtimEY5+xpaGhwcXFxdHQMDQ21PpGbN296eXkFBQUdPXp05syZS5cubfq2W7dujX4mqVSKYRhJklOmTAGAvn37Tps2bcyYMQsXLsQwTC6Xjxgxws/Pz+HXm7QEAgHDMJwTqNEVNGrUqMOHD3O+yokTJxYWFg4ZMoQTn80KplKpxGIxN43YmCDLshiGPT8S69jwgANXa5LSH5sOlwmzKgzPSSFtJlZ0dPTWrVtra2tHjBiBEBoxYsQvv/ySk5PTbEVA48BXWVmpUChqamq4+VqFQrF+/Xocx4ODg0tLS9PT03fv3p2SktI01owZM2bOnGkwGD766KMPP/yw6aiak5MzfPjwgQMHcsbB+fPnx44dSxAE5/dvHHm5H41REEI+Pj46na6+vt5kMiUkJFy+fJnzRjIMQ1FUI29iY2MzM9eV+50AAAHNSURBVDMLCgp69erFpfBcUYpDL2dR+cJuLM12PvZg+a261DozdwdJHx8HjZ56TgpJrFq1yqYIERERxcXFW7Zs2bBhg4eHx8iRIzdt2vTTTz+tXbuWWxSg0WhCQkK46ZfKysoDBw6kpqaWlZXFxcUVFhaePn06ICBg06ZNOI47OzufOHFi27ZtO3fu9PP7bbWan5+fl5fX8uXLAwIClixZwgmSPn36SKXS06dPh4aGhoeHz549OykpKTY2Njk5uaysrH///gkJCSqVqlevXhKJBMOwyMjIqqqq+Pj4tLQ0hUJx6tQpV1fX3r1719bWJicnT5gwISoqSigUFhcX5+bmxsbG1tbWDhw4MDY2dtmyZefOnduzZ49IJKquro6JiWFZ1snJKTAw8PnhlrMAnxEqmxHufKnW/PbJMmd38WAPidHCnrqnWdnfu55ipcTv3B9eMAdpSkrKyZMnIyIiDh8+fOnSJXd3d96nhW2+kxTl6S3EOzsJ/+esonsXl/wS7aex/su6u/HEstl3KhDwi6ofov+Z8qIqg4ViTTTijmRykhC7RwZP8nfgicXjjwb+FnsePLF48MTiwROLBw+eWDx4YvHgicWDB08sHi8C/h+D1oRa6y0riAAAAABJRU5ErkJggg==">
            </li>
          </ul>
        </div>
      </div>
    </nav>

    <script src="https://code.jquery.com/jquery-3.2.1.min.js" integrity="sha256-hwg4gsxgFZhOsEEamdOYGBf13FyQuiTwlAQgxVSNgt4=" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.6.0/Chart.min.js"></script>

    <div class="container-fluid">
      <div class="row">
        <div class="col-sm-3 col-md-2 sidebar">
          <ul class="nav nav-sidebar">
            <li class="active" onclick="changeCCTLD('', this)"><a href="#">All ccTLDs</a></li>
            {{range $cctld := .CCTLDs}}
            <li onclick="changeCCTLD('{{$cctld}}', this)">
              <a href="#"><span class="flag-icon flag-icon-{{$cctld}}"></span> .{{$cctld}}</a>
            </li>
            {{end}}
          </ul>
        </div>
        <div class="col-sm-9 col-sm-offset-3 col-md-10 col-md-offset-2 main">
          <h1 class="page-header">Statistics</h1>
          <h2 class="sub-header" id="title">All Registered Domains</h2>
          <input type="hidden" id="cctld" />

          <canvas id="registered-domains" width="400" height="200"></canvas>

          <script>
            var chart;

            function retrieveStatistics() {
              var url = "/domains/registered";
              if ($("#cctld").val() != "") {
                url += "?cctld=" + $("#cctld").val();
              }

              $.ajax({ url: url })
                .done(function(data) {
                  if (chart) {
                    chart.destroy();
                  }

                  var ctx = document.getElementById("registered-domains");
                  chart = new Chart(ctx, {
                    type: 'line',
                    data: {
                      labels: data.labels,
                      datasets: [{
                        label: "Number of registrations",
                        data: data.data,
                        backgroundColor: "rgba(75,192,192,0.4)",
                        borderColor: "rgba(75,192,192,1)",
                        borderCapStyle: 'butt',
                        borderDash: [],
                        borderDashOffset: 0.0,
                        borderJoinStyle: 'miter',
                        pointBorderColor: "rgba(75,192,192,1)",
                        pointBackgroundColor: "#fff",
                        pointBorderWidth: 1,
                        pointHitRadius: 10,
                        pointHoverRadius: 5,
                        pointHoverBackgroundColor: "rgba(75,192,192,1)",
                        pointHoverBorderColor: "rgba(220,220,220,1)",
                        pointHoverBorderWidth: 2,
                        pointRadius: 1,
                        responsive: true,
                        scaleStartValue: 0,
                        spanGaps: false
                      }]
                    },
                    options: {
                      legend: {
                        display: false
                      }
                    }
                  });
                });
            }

            function changeCCTLD(cctld, caller) {
              if ($("#cctld").val() == cctld) {
                return;
              }

              if (cctld == "") {
                $("#title").text("All Registered Domains");
              } else {
                $("#title").html("<span class='flag-icon flag-icon-" + cctld + "'></span> ." + cctld + " Registered Domains");
              }

              $("#cctld").val(cctld);
              $(".nav li").removeClass("active");
              $(caller).addClass("active");
              retrieveStatistics();
            }

            $(document).ready(function() {
              retrieveStatistics();
            });
          </script>
        </div>
      </div>
    </div>
  </body>
</html>
`

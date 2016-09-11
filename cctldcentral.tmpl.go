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
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/flag-icon-css/2.4.0/css/flag-icon.css" >

    <style>
      /*
       * Base structure
       */

      /* Move down content because we have a fixed navbar that is 50px tall */
      body {
        padding-top: 50px;
      }


      /*
       * Global add-ons
       */

      .sub-header {
        padding-bottom: 10px;
        border-bottom: 1px solid #eee;
      }

      /*
       * Top navigation
       * Hide default border to remove 1px line.
       */
      .navbar-fixed-top {
        border: 0;
      }

      /*
       * Sidebar
       */

      /* Hide for mobile, show later */
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

      /* Sidebar navigation */
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


      /*
       * Main content
       */

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
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
            <span class="sr-only">Toggle navigation</span>
          </button>
          <a class="navbar-brand" href="#">ccTLD Central</a>
        </div>
        <div id="navbar" class="navbar-collapse collapse">
          <ul class="nav navbar-nav navbar-right">
            <li><a href="#">Statistics</a></li>
          </ul>
        </div>
      </div>
    </nav>

    <script src="https://code.jquery.com/jquery-3.1.0.min.js" integrity="sha256-cCueBR6CsyA4/9szpPfrX3s49M9vUU5BgtiJj06wt/s=" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.2.2/Chart.min.js"></script>

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
            function retrieveStatistics() {
              var url = "/domains/registered";
              if ($("#cctld").val() != "") {
                url += "?cctld=" + $("#cctld").val();
              }

              $.ajax({ url: url })
                .done(function(data) {
                  var ctx = document.getElementById("registered-domains");
                  var registeredDomains = new Chart(ctx, {
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

function initLoadBackups() {
    $('.backupsadmin').empty();
    // $.get('http://alphatree.services/getAllBackups', function (data) {
    $.get('http://192.168.0.90/getAllBackups', function (data) {
        console.log(data);
        if (data == null) {
            let foo = "<p>No Backups present</p>"
            $('.backupsadmin').append(foo);
        } else {
            $.each(data, function (key, val) {
                let one = "<p>" + val + "</p>";
                $('.backupsadmin').append(one);
            })
        }
    })
}

function initLoadReviews() {
    $('.review1').empty();
    // $.get('http://alphatree.services/AllApprovedReviews', function (data) {
    $.get('http://192.168.0.90/AllApprovedReviews', function (data) {
        console.log(data);
        if (data == null) {
            let foo = "<p>No Reviews To Process</p>"
            $('.review1').append(foo);
        } else {
            $.each(data, function (key, val) {
                let two = "<div class='rev-card-body'>";
                let three = "<h5 class='rev-card-title'>Review</h5>";
                let four = "<p class='rev-cbod'>" + val.Message + "</p>";
                let five = "<p class='rev-csig'>" + val.Sig + "</p>";
                let six = "</div>";
                let newReview = two + three + four + five + six;
                $('.review1').append(newReview);
            })
        }
    })
}
function initLoadQReviews() {
    $('.reviewadmin').empty();
    // $.get('http://alphatree.services/AllQReviews', function (data) {
    $.get('http://192.168.0.90/AllQReviews', function (data) {
        console.log(data);
        if (data == null) {
            let foo = "<p>No Reviews To Process</p>"
            $('.reviewadmin').append(foo);
        } else {
            $.each(data, function (key, val) {
                let one = "<div class='reviewTop' id='" + val.UUID + "'>";
                let two = "<div class='qborder'>";
                let three = "<h5 >Review" + key + "</h5>";
                let threea = "<p>Name: " + val.Name + "</p>";
                let threeb = "<p>Email: " + val.Email + "</p>"
                let threec = "<p>Unique ID: " + val.UUID + "</p>"
                let four = "<p>Message: " + val.Message + "</p>";
                let five = "<p>Signature: " + val.Sig + "</p>";
                let six = "<button class='delBtn' data-delBtnId='" + val.UUID + "'>Delete Review" + key + "</button></div></div>"
                let newReview = one + two + three + threea + threeb + threec + four + five + six;
                $('.reviewadmin').append(newReview);
            })
        }
    })
}

// function fadeinout() {
//     $('.aniMain').delay(2750).fadeOut(750);
//     $('.mainDiv').delay(3575).fadeIn(1000);
// }

$(document).on("click", "#revInput", function (event) {
    event.preventDefault();
    let name = $("#name").val();
    let email = $("#email").val();
    let message = $("#message").val();
    $.get('atq',
        {
            'name': name,
            'email': email,
            'message': message,
        },
        function (data) {
            console.log("sucess");
        })
    $("#name").val("");
    $("#email").val("");
    $("#message").val("");
})
    .on("click", ".delBtn", function () {
        let rtd = $(this).attr('data-delBtnId');
        let reviewToDeleteUUID = "#" + rtd
        $(reviewToDeleteUUID).empty();
        console.log(reviewToDeleteUUID)
        $.get('/DeleteReview',
            {
                'uuid': rtd,
            },
            function (data) {
                console.log("set Delete sucess");
            })
    })
    .on("click", "#accept", function () {
        $.get('/ProcessQuarintine',
            {},
            function (data) {
                console.log("set Delete sucess");
                let procMess = "<h1>All reviews processed</h1>"
                $('.reviewadmin').empty().append(procMess);
                $.get('/Backup',
                    {},
                    function (data) {
                        console.log("set Backup sucess");
                    })
            })
    })
    .on("click", "#backup", function () {
        $.get('/Backup',
            {},
            function (data) {
                console.log("set Backup sucess");
            })
    })
    .on("click", "#name", function () {
        $("#name").val("")
    })
    .on("click", "#email", function () {
        $("#email").val("")
    })
    .on("click", "#backup", function () {
        initLoadBackups();
    });

$(document).ready(function () {
    // fadeinout();
    initLoadReviews();
    initLoadQReviews();
    initLoadBackups();
});

// document.querySelectorAll(".scene").forEach((elem) => {
// 	const modifier = elem.getAttribute("data-modifier");

// 	basicScroll
// 		.create({
// 			elem: elem,
// 			from: 0,
// 			to: 519,
// 			direct: true,
// 			props: {
// 				"--translateY": {
// 					from: "0",
// 					to: `${10 * modifier}px`
// 				}
// 			}
// 		})
// 		.start();
// });
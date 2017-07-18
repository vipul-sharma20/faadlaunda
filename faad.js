$(document).ready(function() {
    var images = ['https://i.imgflip.com/1rwjm4.jpg', 'https://i.imgflip.com/1rwkc2.jpg',
                  'https://i.imgflip.com/1rwkdv.jpg', 'https://i.imgflip.com/1rwkf9.jpg',
                  'https://i.imgflip.com/1rwkj3.jpg', 'https://i.imgflip.com/1rwkp0.jpg',
                  'https://i.imgflip.com/1rwkrk.jpg'];
    var index = Math.floor((Math.random() * 6) + 0);
    $('#faad_image').attr('src', images[index]);
});

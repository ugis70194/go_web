var rating1 = 0
var rating2 = 0
var jankenRes = 0

axios.defaults.baseURL = 'http://localhost:8000';
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=utf-8';
axios.defaults.headers.post['Access-Control-Allow-Origin'] = '*';

var vm = new Vue({
    el: '#app',
    data: {
        user1: null,
        user2: null,
        rating1: rating1,
        rating2: rating2,
        jankenRes: jankenRes, 
    },
    computed: {
        res() {
            var res = 2
            if(this.rating1 < this.rating2) res = 1
            else if(this.rating1 === this.rating2) res = jankenRes

            if(res === 2) return "勝ち！"
            if(res === 1) return "負け！"　
            if(res === 0) return "引き分け！"
        }
    },
    methods: {
        setInfo() {
            this.rating1 = 1333
            this.rating2 = 1222
            console.log(event.target.tagName)
        },
        getRating() {
            axios.get('http://localhost:8000/user_rating?username=' + this.user1 + '&username=' + this.user2)
            .then( res => ( this.rating1 = res['data']['rating1'],this.rating2 = res['data']['rating2']));
        }
    }
})

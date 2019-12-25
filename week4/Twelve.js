var fs = require("fs")

var data_storage_obj = {
    data: [],
    init: function (filePath) {
        let me = this;
        me["data"] = fs.readFileSync(filePath).toString().replace(/[\W_]+/g, " ").toLowerCase().split(" ");
    },
    words: function () {
        let me = this;
        return me["data"];}
};

var stop_words_obj = {
    stop_words: [],
    init: function () {
        let me = this;
        let a = fs.readFileSync("../stop_words.txt").toString().split(",");
        me["stop_words"] = new Set(a);
    },
    is_stop_word: function (word) {
        let me = this;
        return (word.length === 1) || me["stop_words"].has(word);
    }
};

var word_freqs_obj = {
    freqs: {},
    increment_count: function (word) {
        let me = this;
        if(me["freqs"][word] !== undefined){
            me["freqs"][word] += 1;
        } else {
            me["freqs"][word] = 1;
        }
    },
    sorted: function() {
        let me = this;
        return Object.entries(word_freqs_obj["freqs"]).sort((a,b)=>{return b[1] - a[1];});
    }
};

data_storage_obj["init"](process.argv[2]);
stop_words_obj["init"]();
data_storage_obj["words"]().forEach((item, index, array)=>{
    if(!stop_words_obj["is_stop_word"](item)){
        word_freqs_obj["increment_count"](item);
    }
});
// --- 12.1
// word_freqs_obj["sorted"]().slice(0,25).forEach((value)=>{console.log(value[0], " - ", value[1]);});
// --- 12.2
word_freqs_obj["top25"] = function () {
    let me = this;
    me["sorted"]().slice(0,25).forEach((value)=>{console.log(value[0], " - ", value[1]);});
};
word_freqs_obj["top25"]();

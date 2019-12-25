var fs = require("fs");

var stopwords = new Set(fs.readFileSync("../stop_words.txt").toString().split(","));

function partition(data_str, nlines) {
    let lines = data_str.split("\n");
    let res = [];
    for(let i = 0; i < lines.length; i += nlines) {
        res.push(lines.slice(i, i + nlines).join(" "));
    }
    return res;
}

function split_words(data_str) {
    result = data_str.replace(/[\W_]+/g, " ").toLowerCase().split(" ")
        .filter(word=>(!stopwords.has(word) && word.length > 1))
        .map(word=>[word, 1]);
    return result;
}

function regroup(pairs_list) {
    let mapping = {};
    for (let i in pairs_list){
        for (let pair of pairs_list[i]) {
            if (mapping[pair[0]] !== undefined){
                mapping[pair[0]].push(pair);
            } else {
                mapping[pair[0]] = [pair];
            }
        }
    }
    let res = Object.entries(mapping);
    return res;
}

function count_word(mapping) {
    return [mapping[0], mapping[1].reduce((accum, curV)=>(accum + curV[1]), 0)];
}

function read_file(path) {
    return fs.readFileSync(path).toString();
}
function sort(word_freq) {
    return word_freq.sort((a, b)=>(b[1]-a[1]));
}

var splits = partition(read_file(process.argv[2]), 200).map(split_words);
var splits_per_word = regroup(splits);
var word_freqs = sort(splits_per_word.map(count_word));
word_freqs.slice(0, 25).forEach((value)=>console.log(value[0], " - ", value[1]));
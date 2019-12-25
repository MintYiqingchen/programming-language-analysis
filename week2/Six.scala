import scala.io.Source, scala.collection.mutable
object Main {
    def main(args: Array[String]) :Unit = {
        val stop_words = Source.fromFile("../stop_words.txt").getLines.mkString.split(",").toSet
        val wordmap = mutable.HashMap.empty[String, Int]
        "[a-z]{2,}".r.findAllIn(Source.fromFile(args(0)).getLines.mkString(" ").toLowerCase).filterNot((word:String) => stop_words.contains(word)).foreach((w:String) => if(wordmap.contains(w)) wordmap(w)+=1 else wordmap += (w -> 1))
        wordmap.toList.sortWith(_._2 > _._2).take(25).foreach{ case (w, c) => println(s"$w - $c")}
    }
}
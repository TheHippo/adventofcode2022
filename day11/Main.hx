package;

import haxe.Resource;
import haxe.Log;

using StringTools;
 
enum Operation {
    Plus;
    Multiply;
}

enum Operator {
    Old;
    Value(v:UInt);
}

typedef Throw = {
    to: Int,
    worry: UInt,
}

enum TestResult {
    True;
    False;
}

class Monkey {
    public var index:Int;
    public var items:Array<UInt>;
    public var worryOperation:Operation;
    public var worryOperator:Operator;
    public var test:Int;
    public var throwTo:haxe.ds.Map<TestResult,Int>;
    public var inspections:Int = 0;

    public function new(idx:Int, items:Array<UInt>, worryOperation:Operation, worryOperator:Operator, test:Int, trueMonkey, falseMonkey:Int) {
        this.index = idx;
        this.items =items;
        this.worryOperator = worryOperator;
        this.worryOperation = worryOperation;
        this.test = test;
        this.throwTo = [
            True => trueMonkey,
            False => falseMonkey,
        ];
    }

    public function itemsString():String {
        return items.join(', ');
    }

    public function fetch(t:Throw) {
        trace("fetch", t);
        if (t.to != index) {
            throw "invalid index";
        }
        items.push(t.worry);
    }

    public function getThrow(i:Int):Throw {
        var nw = newWorry(i)%Main.breakFactor;
        trace(nw);
        items.shift();
        var b = switch (nw%test==0) {
            case true: True;
            case false: False;
        }
        return {
            to: throwTo[b],
            worry: nw,
        };
    }

    function newWorry(i:Int):UInt {
        inspections++;
        return switch ([this.worryOperation, this.worryOperator]) {
            case [Plus, Old]: items[i]+items[i];
            case [Multiply, Old]: items[i]*items[i];
            case [Plus, Value(v)]: items[i]+v;
            case [Multiply, Value(v)]: items[i]*v;
        }
    }
}

class Main {

    public static var breakFactor = 1;

    public static function main() {
        var parse = ~/Monkey\s(\d+):\s+Starting items:\s([0-9, ]+)\s+Operation:\snew = old ([*+])\s(old|\d+)\s+Test: divisible by (\d+)\s+If true: throw to monkey (\d+)\s+If false: throw to monkey (\d+)/gm;
        var input = Resource.getString("input");
        var monkeys = new Array<Monkey>();
        while (parse.match(input)) {
            monkeys.push(new Monkey(
                Std.parseInt(parse.matched(1)),
                parseItems(parse.matched(2)),
                parseOperation(parse.matched(3)),
                parseOperator(parse.matched(4)),
                Std.parseInt(parse.matched(5)),
                Std.parseInt(parse.matched(6)),
                Std.parseInt(parse.matched(7))
            ));
            breakFactor *= Std.parseInt(parse.matched(5));
            input = parse.matchedRight();
        }
        // throw breakFactor;

        for (i in 0...10000) {
            // Log.trace(i);
            for (m in monkeys) {
                while (m.items.length > 0) {
                    var t = m.getThrow(0);
                    monkeys[t.to].fetch(t);
                }
            }
            if ((i+1)%1000==0) {
                Log.trace('After round ${i+1}, the monkeys are holding items with these worry levels:');
                for (m in monkeys) {
                    Log.trace('Monkey ${m.index}: ${m.itemsString()}');
                }
            }
        }
        var all = [];
        for (m in monkeys) {
            all.push(m.inspections);
            Log.trace('Monkey ${m.index} inspected items ${m.inspections} times.');
        }
        var highest = twoHighest(all);
        Log.trace(highest[0]*highest[1]);
    }

    static function parseOperator(str:String):Operator {
        var t = str.trim();
        return if (t=="old") {
            Old;
        } else {
            Value(Std.parseInt(str.trim()));
        }
    }

    static function parseOperation(str:String):Operation {
        return switch (str.trim()) {
            case "+": Plus;
            case "*": Multiply;
            case _: throw "invalid operation";
        }
    }

    static function parseItems(str:String):Array<UInt> {
        return [for (i in str.split(',')) Std.parseInt(i.trim())];
    }

    static function twoHighest(nums:Array<Int>):Array<Int> {
        var high1 = 0;
        var high2 = 0;
        for (num in nums) {
          if (num > high1) {
            high2 = high1;
            high1 = num;
          } else if (num > high2) {
            high2 = num;
          }
        }
        return [high1, high2];
     }
}
package;

import haxe.Log;
import haxe.Json;
import haxe.Resource;


enum Value {
    IntValue(v:Int);
    ListValue(v:Array<Value>);
}

class Expression {
    public var values: Array<Value>;

    public function new(line:String) {
        switch (evaluate(Json.parse(line))) {
            case ListValue(v):
                this.values = v;
            case _:
                throw "invalid start value";
        }        
    }

    function evaluate(v:Dynamic):Value {
        switch (Type.typeof(v)) {
            case TClass(Array):
                return ListValue([for(i in 0...v.length) evaluate(v[i])]);
            case TInt:
                return IntValue(v);
            case _:
                throw "unknown type";
        }
    }
}

class Pair {
    var left:Expression;
    var right:Expression;
    public function new(left,right:Expression) {
        this.left = left;
        this.right = right;
    }

    public inline function compare():Bool {
        return this.compareValues(this.left.values, this.right.values);
    }

    private function compareValues(a,b:Array<Value>):Bool {
        var lengthDiff = a.length - b.length;
        var min = a.length > b.length ? b.length : a.length;
        for (i in 0...min) {
            switch ([a[i],b[i]]) {
                case [IntValue(aVal),IntValue(bVal)]:
                    if (aVal > bVal) {
                        return false;
                    }
                case [ListValue(aVal), IntValue(_)]:
                    return compareValues(aVal, [b[i]]);
                case [IntValue(_), ListValue(bVal)]:
                    return compareValues([a[i]],bVal);
                case [ListValue(aVal), ListValue(bVal)]:
                    return compareValues(aVal, bVal);
                case _: 
                    trace(a[i], b[i]);
                    throw "not implemented";
            }
        }
        if (lengthDiff>0) {
            return false;
        }
        return true;
    }
}


class Main {

    public static function main() {
        var exprs = [
            for (line in Resource.getString('input').split('\n')) {
                if (line.length>0) {
                    new Expression(line);
                }
            }
        ];
        var pairs = [
            for (i in 0...exprs.length) {
                if (i%2==0) {
                    continue;
                }
                new Pair(exprs[i-1], exprs[i]);
            }
        ];
        var result = 0;
        for (i in 0...pairs.length) {
            if (pairs[i].compare()) {
                trace(i);
                result += i+1;
            }
        }
        Log.trace(result);
    }
}
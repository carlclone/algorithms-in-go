package bit_manipulation

func hammingWeight(num uint32) int {
	var (
		res int
	)
	for num != 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

func hammingWeight(num uint32) int {
	res := 0
	mask := uint32(1)
	for i := 1; i <= 32; i++ {
		if num&mask == 1 {
			res += 1
		}
		mask = mask << 1
	}
	return res
}

//异或,相同的为0 , 不同的为1

//https://www.jianshu.com/p/4e73512c03b8
// mask 掩码 , 面具 , 掩藏 , 用于屏蔽某些二进制位 , 与运算 (我都快不认识掩这个字了)
// 子网掩码 , 提取子网 ,   php or redis字典掩码 , 映射到范围内 , 类似取余

//位运算应用口诀
//清零取反要用与，某位置一可用或
//若要取反和交换，轻 轻松松用异或
//
//移位运算
//要点 1 它们都是双目运算符，两个运算分量都是整形，结果也是整形。
//     2 "<<" 左移：右边空出的位上补0，左边的位将从字头挤掉，其值相当于乘2。
//     3 ">>"右移：右边的位被挤掉。对于左边移出的空位，如果是正数则空位补0，若为负数，可能补0或补1，这取决于所用的计算机系统。
//     4 ">>>"运算符，右边的位被挤掉，对于左边移出的空位一概补上0。
//
//位运算符的应用 (源操作数s 掩码mask)
//(1) 按位与-- &
//1 清零特定位 (mask中特定位置0，其它位为1，s=s&mask)
//2 取某数中指定位 (mask中特定位置1，其它位为0，s=s&mask)
//(2) 按位或-- |
//    常用来将源操作数某些位置1，其它位不变。 (mask中特定位置1，其它位为0 s=s|mask)
//(3) 位异或-- ^
//1 使特定位的值取反 (mask中特定位置1，其它位为0 s=s^mask)
//2 不引入第三变量，交换两个变量的值 (设 a=a1,b=b1)
//    目标           操作              操作后状态
//a=a1^b1         a=a^b              a=a1^b1,b=b1
//b=a1^b1^b1      b=a^b              a=a1^b1,b=a1
//a=b1^a1^a1      a=a^b              a=b1,b=a1
//
//二进制补码运算公式：
//-x = ~x + 1 = ~(x-1)
//~x = -x-1
//-(~x) = x+1
//~(-x) = x-1
//x+y = x - ~y - 1 = (x|y)+(x&y)
//x-y = x + ~y + 1 = (x|~y)-(~x&y)
//x^y = (x|y)-(x&y)
//x|y = (x&~y)+y
//x&y = (~x|y)-~x
//x==y:    ~(x-y|y-x)
//x!=y:    x-y|y-x
//x< y:    (x-y)^((x^y)&((x-y)^x))
//x<=y:    (x|~y)&((x^y)|~(y-x))
//x< y:    (~x&y)|((~x|y)&(x-y))//无符号x,y比较
//x<=y:    (~x|y)&((x^y)|~(y-x))//无符号x,y比较
//
//应用举例
//(1) 判断int型变量a是奇数还是偶数
//a&1   = 0 偶数
//       a&1 =   1 奇数
//(2) 取int型变量a的第k位 (k=0,1,2……sizeof(int))，即a>>k&1
//(3) 将int型变量a的第k位清0，即a=a&~(1<<k)
//(4) 将int型变量a的第k位置1， 即a=a|(1<<k)
//(5) int型变量循环左移k次，即a=a<<k|a>>16-k   (设sizeof(int)=16)
//(6) int型变量a循环右移k次，即a=a>>k|a<<16-k   (设sizeof(int)=16)
//(7) 整数的平均值
//对于两个整数x,y，如果用 (x+y)/2 求平均值，会产生溢出，因为 x+y 可能会大于INT_MAX，但是我们知道它们的平均值是肯定不会溢出的，我们用如下算法：
//int average(int x, int y)   //返回X,Y 的平均值
//{
//     return (x&y)+((x^y)>>1);
//}
//(8)判断一个整数是不是2的幂,对于一个数 x >= 0，判断他是不是2的幂
//boolean power2(int x)
//{
//    return ((x&(x-1))==0)&&(x!=0)；
//}
//(9)不用 temp交换两个整数
//void swap(int x , int y)
//{
//    x ^= y;
//    y ^= x;
//    x ^= y;
//}
//(10) 计算绝对值
//int abs( int x )
//{
//int y ;
//y = x >> 31 ;
//return (x^y)-y ;        //or: (x+y)^y
//}
//(11) 取模运算转化成位运算 (在不产生溢出的情况下)
//         a % (2^n) 等价于 a & (2^n - 1)
//(12)乘法运算转化成位运算 (在不产生溢出的情况下)
//         a * (2^n) 等价于 a<< n
//(13)除法运算转化成位运算 (在不产生溢出的情况下)
//         a / (2^n) 等价于 a>> n
//        例: 12/8 == 12>>3
//(14) a % 2 等价于 a & 1
//(15) if (x == a) x= b;
//　 　          else x= a;
//　　      等价于 x= a ^ b ^ x;
//(16) x 的 相反数 表示为 (~x+1)

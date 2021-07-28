package Java.se;

import java.util.Arrays;
import java.util.List;

// java8的新特性：方法引用
// 四种使用方式：1.类名::静态方法名，2.对象::实例方法名，3.类名::实例方法名(即类plus非静态方法)，4.类名::new
public class MethodReferenceDemo {
    public static void main(String[] args) {
        // 首先使用lambda表达式
        Student student1 = new Student("zhangsan",60);
        Student student2 = new Student("lisi",70);
        Student student3 = new Student("wangwu",80);
        Student student4 = new Student("zhaoliu",90);
        List<Student> students = Arrays.asList(student1,student2,student3,student4);
        students.sort((o1, o2) -> o1.getScore() - o2.getScore());
        students.forEach(student -> System.out.println(student.getScore()));

        // sort方法接收一个Comparator函数式接口，接口中唯一的抽象方法compare接收两个参数返回一个int类型值，下方是Comparator接口定义
        // @FunctionalInterface
        // public interface Comparator<T> {
        //     int compare(T o1, T o2);
        // }
        // Student类中定义的compareStudentByScore静态方法是：
        //public static int compareStudentByScore(Student student1,Student student2){
        //     return student1.getScore() - student2.getScore();
        // }
        // 签名一致，即除了方法名外，接受的参数和返回值的类型相同，便可以用引用方法来替换
        // 1.类名::静态方法名 方法引用替换lambda表达式
        students.sort(Student::compareStudentByScore);

        // 2.对象::实例方法名
        StudentComparator studentComparator = new StudentComparator();
        students.sort(studentComparator::compareStudentByScore);

        // 3.类名::实例方法名，比较特殊
        // 类名::实例方法名 方法引用时，一定是lambda表达式所接收的第一个参数来调用实例方法，如果lambda表达式接收多个参数，其余的参数作为方法的参数传递进去
        // students.sort((o1, o2) -> o1.getScore() - o2.getScore()); 所以这里局势o1来调用Student::compareByScore，然后o2作为这个方法的参数
        students.sort(Student::compareByScore);

        // 4.类名::new, 构造引用，和前两种类似只要符合lambda表达式的定义即可
        
    }

}

class Student {
    private String name;
    private int score;

    public Student(){

    }

    public Student(String name, int score){
        this.name = name;
        this.score = score;
    }

    public String getName(){
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    public int getScore() {
        return score;
    }
    public void setScore(int score) {
        this.score = score;
    }

    public static int compareStudentByScore(Student student1,Student student2){
        return student1.getScore() - student2.getScore();
    }

    public static int compareStudentByName(Student student1,Student student2){
        return student1.getName().compareToIgnoreCase(student2.getName());
    }

    public int compareByScore(Student student){
        return this.getScore() - student.getScore();
    }

}

class StudentComparator {
    public int compareStudentByScore(Student student1,Student student2){
        return student2.getScore() - student1.getScore();
    }
}

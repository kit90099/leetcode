public class main {
    public static void main(String... args){
        ExamRoom er = new ExamRoom(10);
        System.out.println(er.seat());
        System.out.println(er.seat());
        System.out.println(er.seat());
        System.out.println(er.seat());
        er.leave(4);
        System.out.println(er.seat());
    }
}

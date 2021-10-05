import java.util.*;
class selection
//select a cell and compare it with all other cells
{
    public static void main(String args[])
    {
        Scanner d=new Scanner(System.in);
        int i,j,t,k;
        int a[]=new int[10];
        for(i=0;i<10;i++)
            a[i]=d.nextInt();
        for(i=0;i<9;i++)//0 -> n-1
        {
            for(j=i+1;j<10;j++)//i+1 ->n
            {
                if(a[i]>a[j])//a[i] , a[j]
                {
                    t=a[i];
                    a[i]=a[j];
                    a[j]=t;
                }
            }
            for(k=0;k<10;k++)
            System.out.print(a[k]+" ");
            System.out.println();
        }
        for(i=0;i<10;i++)
            System.out.print(a[i]+" ");
    }
}

import java.util.*;
class bubble
//select a cell and compare it with next cell
{
    public static void main(String args[])
    {
        Scanner d=new Scanner(System.in);
        int i,j,t,k,min,p;
        int a[]=new int[10];
        for(i=0;i<10;i++)
            a[i]=d.nextInt();
        for(i=0;i<10;i++)//0 ->n
        {
            for(j=0;j<9-i;j++)//0 ->n-1-i
            {
                if(a[j]>a[j+1])//next element
                {
                    t=a[j];
                    a[j]=a[j+1];
                    a[j+1]=t;
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

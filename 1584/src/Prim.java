import java.util.List;
import java.util.PriorityQueue;

class Prim {
    private PriorityQueue<int[]> pq;
    private boolean[] inMst;
    private int weightSum = 0;
    private List<int[]>[] graph;

    public Prim(List<int[]>[] graph){
        this.graph = graph;
        this.pq = new PriorityQueue<>((a,b)->{
            return a[2] - b[2];
        });

        int n = graph.length;
        this.inMst = new boolean[n];

        inMst[0] = true;
        cut(0);
        while(!pq.isEmpty()){
            int[] edge = pq.poll();
            int to = edge[1];
            int weight = edge[2];
            if (inMst[to]){
                continue;
            }

            weightSum += weight;
            inMst[to] = true;

            cut(to);
        }
    }

    private void cut(int s){
        for (int[] edge: graph[s]){
            int to = edge[1];
            if (!inMst[to]){
                pq.offer(edge);
            }
        }
    }

    public int getWeightSum(){
        return weightSum;
    }

    public boolean isAllConnected(){
        for(boolean point:inMst){
            if(!point){
                return false;
            }
        }
        return true;
    }
}
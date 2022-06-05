import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

public class Solution {
    public int minimumCost(int n, int[][] connections) {
        List<int[]>[] graph = buildGraph(n, connections);

        Prim prim = new Prim(graph);

        if(!prim.isAllConnected()){
            return -1;
        }

        return prim.getWeightSum();
    }

    private List<int[]>[] buildGraph(int n, int[][] connections) {
        List<int[]>[] graph = new LinkedList[n];
        for (int i = 0; i < n; i++) {
            graph[i] = new LinkedList<>();
        }

        for (int[] conn : connections) {
            int u = conn[0] - 1;
            int v = conn[1] - 1;
            int weight = conn[2];

            graph[u].add(new int[]{u,v,weight});
            graph[v].add(new int[]{v,u,weight});
        }

        return graph;
    }
}

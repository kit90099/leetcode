import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

class Solution {
    public int minCostConnectPoints(int[][] points) {
        List<int[]>[] graph = buildGraph(points.length, points);

        Prim prim = new Prim(graph);

        return prim.getWeightSum();
    }

    private List<int[]>[] buildGraph(int n, int[][] connections) {
        List<int[]>[] graph = new LinkedList[n];
        for (int i = 0; i < n; i++) {
            graph[i] = new LinkedList<>();
        }

        for (int i = 0; i < connections.length; i++) {
            for(int j=i+1;j<connections.length;j++){
                int weight = Math.abs(connections[i][0]-connections[j][0])+Math.abs(connections[i][1]-connections[j][1]);

                graph[i].add(new int[]{i, j, weight});
                graph[j].add(new int[]{j, i, weight});
            }
        }

        return graph;
    }
}
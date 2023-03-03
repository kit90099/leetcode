import java.util.*

class Solution {
    private lateinit var indegree: Array<Int>

    fun canFinish(numCourses: Int, prerequisites: Array<IntArray>): Boolean {
        val graph = buildGraph(numCourses, prerequisites)

        indegree = Array<Int>(numCourses){0}
        for (edge in prerequisites) {
            indegree[edge[0]]++
        }

        val queue: Queue<Int> = LinkedList()
        for ((i, value) in indegree.withIndex()) {
            if (value == 0) {
                queue.offer(i)
            }
        }

        var count = 0
        while (queue.isNotEmpty()){
            val curr = queue.poll()
            count++
            for (next in graph[curr]){
                indegree[next]--
                if (indegree[next] == 0){
                    queue.offer(next)
                }
            }
        }

        return count == numCourses
    }

    private fun buildGraph(numCourses: Int, prerequisites: Array<IntArray>): Array<MutableList<Int>> {
        val graph = Array<MutableList<Int>>(numCourses) { mutableListOf() };

        for (prerequisite in prerequisites) {
            graph[prerequisite[1]].add(prerequisite[0])
        }

        return graph
    }
}
import java.util.*

class Solution {
    private lateinit var indegree: Array<Int>

    fun findOrder(numCourses: Int, prerequisites: Array<IntArray>): IntArray {
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
        val res = IntArray(numCourses)
        while (queue.isNotEmpty()){
            val curr = queue.poll()
            res[count] = curr
            count++
            for (next in graph[curr]){
                indegree[next]--
                if (indegree[next] == 0){
                    queue.offer(next)
                }
            }
        }

        if (count != numCourses){
            return IntArray(0)
        }
        return res
    }

    private fun buildGraph(numCourses: Int, prerequisites: Array<IntArray>): Array<MutableList<Int>> {
        val graph = Array<MutableList<Int>>(numCourses) { mutableListOf() };

        for (prerequisite in prerequisites) {
            graph[prerequisite[1]].add(prerequisite[0])
        }

        return graph
    }
}
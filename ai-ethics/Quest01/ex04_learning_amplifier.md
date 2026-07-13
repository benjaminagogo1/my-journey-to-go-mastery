# Phase 1: Build Foundation (you first)

Before building the project, 

- I engaged in active reading by studying
 
- documentation and research papers on wireless routing protocols.
 
Afterward, 

- I tested my understanding by explaining the concepts to myself 

without referring back to the materials.

## Small Network Design (8 Devices)

## Scenario:
### A small smart-building floor with environmental monitoring.

_Devices (8 total):_

-  6 temperature/humidity sensors (static, battery-powered)

- 1 central gateway node (connected to cloud/server)

- 1 maintenance laptop (occasionally connected, mostly static)

## Network Characteristics

- Nodes are static

- Low data rate (periodic sensor readings)

- Battery-powered sensors → energy efficiency is critical

- Multi-hop communication required (some sensors out of direct gateway range)

## Routing Protocol Chosen

Routing Protocol for Low-Power and Lossy Networks (RPL)

## Justification

- Designed specifically for low-power and lossy wireless networks.

- Supports energy-aware routing.

- Forms a Directed Acyclic Graph (DAG) toward the gateway.

- Suitable for mostly static IoT environments.

- Scales better than simple reactive protocols in sensor networks.


# Phase 2: Strategic AI Use


After building my foundational understanding, I moved to strategic AI
use. 

Instead of asking the AI to explain everything from scratch, 

I used it to test my understanding through targeted questions, explore edge cases, and validate my design decisions.

## Edge cases explored with targeted questions.

## Question 1: If one of your 5 IoT devices suddenly fails, how will the others maintain connectivity?


- If a node fails, its neighbors detect the link failure.

- They remove the failed node from their parent/neighbor list.

- Each affected node selects a new parent (or next-hop) from the available neighbors.

>In RPL (used for IoT sensors), this is handled automatically via DIO/DAO messages and rank recalculation.

# Result:

- Network self-heals.

- Routes adapt dynamically.

- Some temporary delays may occur, but connectivity is restored.


## Question 2:
## If two IoT nodes fail simultaneously and the remaining nodes have only one possible parent, what potential problem could occur?

If the remaining nodes have only one possible parent:

- They will struggle to maintain connectivity if that single parent becomes overloaded or fails.

- Traffic may congest through that one parent, increasing latency and packet loss.

- Some nodes may experience temporary isolation if the parent cannot handle all child nodes.

- Energy consumption for that parent increases faster, potentially causing early battery depletion (for IoT devices).

>In short: The network becomes fragile and less resilient because it lacks redundancy.


## Question 3:
## For emergency vehicles using the Ad hoc On-Demand Distance Vector routing protocol, what happens if all vehicles in the same area try to find new routes at the same time? What kind of network problem could occur?

When all vehicles try to find new routes simultaneously, each sends a route request message.

These messages propagate through the network, creating a flood of control traffic.

This leads to:

- Network congestion

- Collisions and packet loss

- Delays in route discovery

**This problem is commonly called a broadcast storm**.

>In short: Too many simultaneous route discoveries overwhelm the network, reducing performance.

## For traffic lights using a proactive mesh protocol, what happens if too many links fail at once? How can the network maintain reliability?

If too many links fail at once in a traffic light network using a proactive mesh protocol:

Immediate effect:

 - Some traffic lights may lose direct communication with their neighbors.

- Control messages from the primary controller may not reach all nodes.

### How the network maintains reliability:

- The protocol continuously maintains routing tables, so alternative paths are already known.

- Traffic lights automatically reroute messages through other available neighbors.

- Backup or secondary controllers can take over if the primary controller becomes unreachable.

- This ensures low-latency and continuous synchronization even during multiple link failures.

>In short: Proactive mesh protocols provide redundancy and multiple paths, allowing the network to recover quickly from multiple simultaneous link failures.
<br><br>


## My reasoning:

Before using AI, I first answered the edge-case questions
 
based on my own understanding. For example, when asked what

would happen if IoT nodes failed, I reasoned that the

remaining nodes would remove the failed node and reattach to 

available neighbors. I also identified that simultaneous route 

discovery in emergency vehicles would cause flooding. After 

presenting my answers, I used AI to validate whether my 

reasoning aligned with how the routing protocols actually 

function. The AI feedback confirmed most of my assumptions and 

added deeper technical explanations about congestion, routing 

instability, and control overhead. This process helped me 

verify that my understanding was technically sound rather than assumed.

<br><br>

# Phase 3: Real Application

My smart-city network design consist of three components which are:

- IoT Sensor Network (1,000 Nodes)

- Traffic Light Network (50 Nodes)

- Emergency Vehicles (10 Nodes)


## IoT Sensor Network (1,000 Nodes)

I designed a smart-city network consisting of three different network behaviors.

>The first component includes 1,000 IoT sensors that are:

- Static (fixed position)

- Battery-powered

- Transmitting small, periodic data packets

>Since the network requires:

- Static node support

- Scalability to a large number of devices

- Low-power operation

>The most appropriate routing protocol is **Routing Protocol for Low-Power and Lossy Networks (RPL).**

## Justification

I selected RPL because:

- 1. It is specifically designed for low-power and lossy networks (LLNs).

- 2. It constructs a Destination-Oriented Directed Acyclic Graph (DODAG) toward a central gateway.

- 3. It supports energy-aware routing metrics, helping to extend battery life.

- 3. It scales efficiently to hundreds or thousands of nodes.

- 4. It is optimized for periodic sensor data transmission.
<br><br>


## Traffic Light Network (50 Nodes)

>The second component consists of 50 traffic lights, which are:

- Fixed infrastructure (static position)

- Powered by electricity (not battery-constrained)

- Requiring reliable synchronization

- Requiring moderate to low latency

- perating in a relatively stable topology

>Since this network requires:

- 1. Stable and continuously maintained routes

- 2. Low latency for control message exchange

- 3. Fast availability of routing information

- 4. Reliable communication between intersections

>The most appropriate routing protocol is **Optimized Link State Routing Protocol (OLSR).**

## Justification

>I selected OLSR because:

- It is a proactive routing protocol, meaning it maintains routing tables continuously.

- Routes are immediately available when needed, reducing latency.

- It performs well in relatively stable network topologies.

- It uses Multipoint Relays (MPRs) to reduce control overhead while maintaining efficiency.

- Since traffic lights are powered by electricity, energy consumption is not a primary constraint


## Emergency Vehicles (10 Nodes)

>The third component consists of  10 Emergency Vehicles

These nodes are:

- Highly mobile

- Frequently changing topology

- Require very fast route adaptation

- Possibly vehicle-to-vehicle (V2V) communication

- Require low latency for priority signaling

### Best Fit Protocol:
**Ad hoc On-Demand Distance Vector (AODV)**

## JUSTIFICATION:

>I chose Ad hoc On-Demand Distance Vector (AODV) because:

- It is a reactive routing protocol, meaning routes are created only when needed.

- It adapts quickly to topology changes, which is critical for moving vehicles.

- It reduces unnecessary routing overhead compared to proactive protocols.

- It performs well in high-mobility wireless ad hoc networks (MANETs).

- It supports fast route discovery when a vehicle enters a new area



# Failure Points:

The following are the failure points.

- 1. Parent node failure

- 2. Energy depletion

- 3. Network partitioning

- 4. High mobility causing flooding




# AI refine


What Happens When a Parent Node Fails?

In RPL:

- Nodes form a DODAG (Destination-Oriented Directed Acyclic Graph) toward the gateway.

- Each sensor chooses a preferred parent (next hop toward the root).

- Nodes usually maintain backup parent options.

## 1. If a parent node fails:

- The child node detects link failure (e.g., no acknowledgment, signal loss).

- It removes that parent from its routing table.

- It selects another available parent from its neighbor list.

- It updates its rank and reattaches to the DODAG.

>So recovery is not manual — it is automatic and distributed.

## If No Alternative Parent Exists

- If a node has no other reachable neighbor:

- It becomes temporarily disconnected.

- It listens for new DODAG Information Object (DIO) messages.

- When another nearby node joins or signal conditions improve, it reconnects.

## Why This Works

**RPL is designed for:**

- Unstable wireless links

- Battery-powered nodes

- Gradual topology changes

>It constantly exchanges lightweight control messages that allow topology repair without flooding the entire network.



## 2. What Happens in AODV During High Mobility

When multiple highly mobile nodes (emergency vehicles) need routes simultaneously:

- 1. Route Discovery Flooding

Each node broadcasts a route request (RREQ) to find a path.

If many nodes do this at the same time, the network experiences flooding.

- 2. Broadcast Storm / Congestion

Many RREQs collide, causing packet loss.

Links may get congested.

Delays increase for all route discoveries.

- 3.Temporary Route Instability

Routes break often because nodes move.

Data packets may be dropped until a new route is found.


## 3. If too many nodes change parents frequently, what problem might occur in the network?


If too many nodes change parents frequently (for example in an RPL DODAG), the network can suffer from several serious problems:

1. Routing Instability (Route Flapping)

Nodes continuously detach and reattach to different parents.
This creates constant topology changes, making routes unstable.

## Result:

Packets may follow inconsistent paths

Increased packet loss

2. Control Overhead Explosion

Every parent change triggers control messages (DIO, DAO updates in RPL).

## Result:

Excessive signaling traffic

Wasted bandwidth

Congestion


## 4. Handling Single-Point-of-Failure in OLSR for Traffic Lights

Problem:

OLSR is proactive and relies on nodes maintaining routing tables.

If a central controller (gateway or main traffic server) fails, all nodes depending on it may lose coordination.

## Solution
Redundancy / Backup Controllers:

- Deploy backup controllers (secondary nodes) that can take over automatically.

- Use controller election protocols or pre-configured priorities to switch seamlessly.

- Each traffic light maintains routes to both primary and secondary controllers.


<br><br>
<br>

<p style = "text-align: center; font-weight: bold; font-size: x-large; text-transform: capitalize">This is a conceptual model of my smart city network design.</p>
<p align="center">
  <img src="image/My-Smart-City-Design-Diagram.png" alt="Centered Image" width="700">
</p>

# Reflection:

## % Human Judgment vs. AI Contribution

The core architectural decisions were based on my understanding of routing protocol categories (proactive, reactive, and hierarchical) and matching them to node behavior (static, mobile, low-power, infrastructure-based). I identified that IoT sensors required low-energy routing, traffic lights required low-latency stability, and emergency vehicles required fast route adaptation.

**AI assisted mainly in**:

- Refining my technical justifications

- Improving precision of terminology

- Confirming that protocol choices aligned with standard networking principles

- The structural thinking and system segmentation were my decisions.

## Could you Defend Decisions Without AI?

Yes, at a conceptual level.

**I understand:**

- Why low-power networks require energy-aware routing.

- Why stable infrastructure benefits from proactive routing.

- Why mobility requires adaptive routing mechanisms.

However, AI helped strengthen my confidence and improved the clarity of the explanations. Without AI, my reasoning would be less polished but still logically defensible.

## What Will you Still Remember in 6 Months?

>I will remember:

- Routing protocol selection depends on node behavior, not popularity.

- Network design is about trade-offs (energy vs latency vs mobility vs scalability).

- One protocol rarely fits all components in a heterogeneous system.

- Mobility dramatically changes routing requirements.

These are conceptual principles that go beyond memorization.

## Did AI Make you Sharper or Think For you?

AI made me sharper.

It did not generate the design from nothing. Instead, it:

- Challenged my assumptions

- Forced clearer justification

- Helped refine technical accuracy

- Exposed gaps in my understanding

**The thinking process remained mine. AI functioned as a technical reviewer, not a replacement for reasoning.**

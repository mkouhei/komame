==============
 Design notes
==============

Design
======

discovery
---------

* The nodes that are enable to log into shell are enable to have two discovery way.
  Those are the installing multicast DNS(mDNS) agent and the polling from the discovery daemon. (calls ``the agent node``.)

* The nodes that are disable to log into shell have only one discovery way
  that is the polling from the discovery daemon(calls ``the discoverd``. Those nodes are named from MAC address with `MA-L Public Listing <http://standards.ieee.org/develop/regauth/oui/public.html>`_. (calls ``the agentless node``.)

DNS
---

* The agent nodes are enabled to lookup DNS with mDNS agent only.
* The agentless nodes need the authoritative server for the LAN.
* Use the metadata acquired from agent and discoverd to `PowerDNS API <https://doc.powerdns.com/md/httpapi/README/>`_ that is the built-in REST/JSON API of the PowerDNS (authoritative server >= 3.4).

Data flow
=========

.. blockdiag:: /_static/komame.diag

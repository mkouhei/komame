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

Data flow
=========

.. blockdiag:: /_static/komame.diag

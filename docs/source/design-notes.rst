==============
 Design notes
==============

Design
======

discovery
---------

* The nodes that are enable to log into shell are enable to have two discovery way.
  Those are the installing multicast DNS agent and the polling from the discovery daemon.
* The nodes that are disable to log into shell have only one discovery way
  that is the polling from the discovery daemon.

Data flow
=========

.. blockdiag:: /_static/komame.diag
